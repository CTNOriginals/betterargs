package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/CTNOriginals/betterargs"
)

// run program.exe --file file/path/file.exe
var argOptions = betterargs.MArgs{
	"help": {
		Flags:       betterargs.Flags{"--help", "-H"},
		Description: "Display a list of possible arguments along with their description",
	},
	"file": {
		Flags:       betterargs.Flags{"--file", "-F"},
		Description: "The file to do stuff with",
		Inputs: betterargs.MInputs{
			"target-file": betterargs.Input{
				Description: "The target file path",
				Required:    true,
				Validator: func(arg string) bool {
					return strings.HasPrefix(arg, "/")
				},
			},
		},
	},
}

var testArgs = []string{"C:\\path\\to\\file\\betterargs.exe",
	"--help",
	"--file", "/path/to/file.ext",
}

func main() {
	fmt.Printf("\n\n---- Start %s ----\n", time.Now().Format(time.TimeOnly))

	fmt.Println(betterargs.ParseArguments(testArgs, argOptions))
}
