package cli

import (
	"os"

	"github.com/tdewolff/argp"
)

func Execute() {
	cmd := argp.New("Unity Package Extractor")
	cmd.AddCmd(&ExtractCommand{}, "x", "extract unity package")
	cmd.AddCmd(&SusCommand{}, "s", "examine sus level of unity package")

	if len(os.Args) <= 1 {
		cmd.PrintHelp()
		os.Exit(0)
	}

	cmd.Parse()
}
