package upkg

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ScanForSus(src string) ([]SuspectFile, error) {
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

	var suspects []SuspectFile
	pathMap := make(map[string]string)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if strings.HasSuffix(hdr.Name, "/pathname") {
			content, err := io.ReadAll(tr)
			if err != nil {
				return nil, err
			}
			dir := filepath.Dir(hdr.Name)
			pathMap[dir] = strings.TrimSpace(string(content))
		}
	}

	for _, path := range pathMap {
		ext := strings.ToLower(filepath.Ext(path))

		var t SuspectFileType

		switch ext {
		case ".dll":
			t = SuspectDLL
		case ".cs":
			t = SuspectCS
		default:
			continue
		}

		suspects = append(suspects, SuspectFile{
			Path: path,
			Type: t,
		})

	}

	return suspects, nil
}
