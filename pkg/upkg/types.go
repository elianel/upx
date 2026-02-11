package upkg

type SuspectFileType string

const (
	SuspectDLL SuspectFileType = "dll"
	SuspectCS  SuspectFileType = "cs"
)

type SuspectFile struct {
	Path string
	Type SuspectFileType
}

type ExtractResult struct {
	Extracted []string
	Skipped   []string
}
