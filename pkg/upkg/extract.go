package upkg

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Extract(src string, dst string) (*ExtractResult, error) {
	if src == "" {
		return nil, errors.New("source path required")
	}
	if dst == "" {
		return nil, errors.New("destination path required")
	}

	if err := os.MkdirAll(dst, 0755); err != nil {
		return nil, err
	}

	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	tmpDir, err := os.MkdirTemp("", "unitypkg-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	// Step 1: unpack full tar into temp
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		target := filepath.Join(tmpDir, hdr.Name)

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return nil, err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return nil, err
			}
			outFile, err := os.Create(target)
			if err != nil {
				return nil, err
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return nil, err
			}
			outFile.Close()
		}
	}

	result := &ExtractResult{}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		entryDir := filepath.Join(tmpDir, entry.Name())
		pathFile := filepath.Join(entryDir, "pathname")
		assetFile := filepath.Join(entryDir, "asset")

		pathBytes, err := os.ReadFile(pathFile)
		if err != nil {
			result.Skipped = append(result.Skipped, entry.Name())
			continue
		}

		relPath := strings.TrimSpace(string(pathBytes))
		finalPath := filepath.Join(dst, relPath)

		// prevent path traversal
		if !strings.HasPrefix(filepath.Clean(finalPath), filepath.Clean(dst)) {
			result.Skipped = append(result.Skipped, relPath)
			continue
		}

		if _, err := os.Stat(assetFile); err != nil {
			result.Skipped = append(result.Skipped, relPath)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(finalPath), 0755); err != nil {
			return nil, err
		}

		if err := moveFile(assetFile, finalPath); err != nil {
			return nil, err
		}

		result.Extracted = append(result.Extracted, relPath)
	}

	return result, nil
}
func moveFile(src, dst string) error {
	// Try rename first (fast path)
	if err := os.Rename(src, dst); err == nil {
		return nil
	} else {
		// If not cross-device, return error immediately
		if !isCrossDeviceErr(err) {
			return err
		}
	}

	// Fallback: copy + remove
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	if err := out.Sync(); err != nil {
		return err
	}

	return os.Remove(src)
}

func isCrossDeviceErr(err error) bool {
	// linux
	return strings.Contains(err.Error(), "invalid cross-device link")
}
