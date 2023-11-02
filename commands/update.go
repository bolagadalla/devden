package commands

import (
	"flag"
	"os"
)

func HandleUpdate(update *flag.FlagSet, config *string, name *string, desc *string) {
	update.Parse(os.Args[2:])
}
