package cli

import (
	"log"

	"github.com/elianel/upx/pkg/upkg"
)

type ExtractCommand struct {
	Src, Dst string
}

func (cmd *ExtractCommand) Run() error {
	log.Printf("[UPX] Extracting src=%s -> dst=%s", cmd.Src, cmd.Dst)

	result, err := upkg.Extract(cmd.Src, cmd.Dst)
	if err != nil {
		return err
	}

	for _, f := range result.Extracted {
		log.Printf("[UPX] Extracted: %s", f)
	}

	// for _, f := range result.Skipped {
	// 	log.Printf("[UPX] Skipped: %s", f)
	// }

	return nil
}
