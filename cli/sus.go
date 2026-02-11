package cli

import (
	"log"

	"github.com/elianel/upx/pkg/upkg"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

type SusCommand struct {
	Src string
}

func (cmd *SusCommand) Run() error {
	log.Printf("[UPX] Examining src=%s", cmd.Src)

	suspects, err := upkg.ScanForSus(cmd.Src)
	if err != nil {
		return err
	}

	for _, s := range suspects {
		log.Printf("[UPX] %sSus File Detected%s: %s (%s)",
			colorRed,
			colorReset,
			s.Path,
			s.Type,
		)
	}

	if len(suspects) == 0 {
		log.Printf("[UPX] Result=%sSafe%s",
			colorGreen,
			colorReset,
		)
		return nil
	}

	log.Printf("[UPX] Result=%sSUS%s",
		colorRed,
		colorReset,
	)
	return nil
}
