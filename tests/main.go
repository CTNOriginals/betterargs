package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/CTNOriginals/betterargs"
)

// run program.exe --file file/path/file.exe
var argOptions = betterargs.MFlags{
	"help": {
		Aliases:     []string{"-H"},
		Description: "Display a list of possible arguments along with their description",
	},
	"file": {
		Description: "The file to do stuff with",
		Inputs: betterargs.InputOrder{
			{
				Name:        "target-dir",
				Description: "The target directory path",
				Validator: func(arg string) bool {
					return strings.HasSuffix(arg, "/")
				},
			},
			{
				Name:        "target-file",
				Description: "The target file path",
				Required:    true,
				Validator: func(arg string) bool {
					return !strings.HasSuffix(arg, "/")
				},
			},
		},
	},
}

var testArgs = []string{"C:\\path\\to\\file\\betterargs.exe",
	"--file", "/path/to/dir/", "path/to/file.ext",
	"--file", "path/to/file.ext",
	"--help",
}

func main() {
	fmt.Printf("\n\n---- Start %s ----\n", time.Now().Format(time.TimeOnly))

	fmt.Println(betterargs.ParseArguments(testArgs, argOptions))
}
