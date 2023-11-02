package commands

import (
	"flag"
	"os"
)

func HandleCreate(create *flag.FlagSet, pn *string) {
	create.Parse(os.Args[2:])
}
