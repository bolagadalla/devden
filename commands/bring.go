package commands

import (
	"flag"
	"os"
)

func HandleBring(bring *flag.FlagSet, config *bool) {
	bring.Parse(os.Args[2:])
}
