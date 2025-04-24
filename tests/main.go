package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/CTNOriginals/betterargs"
)

// run program.exe --file file/path/file.exe
var argOptions = betterargs.MFlags{
	// "help": {
	// 	Aliases:     []string{"-H"},
	// 	Description: "Display a list of possible arguments along with their description",
	// },
	"files": {
		Description: "The file to do stuff with",
		Inputs: betterargs.InputOrder{
			{
				Name:          "search-directories",
				Description:   "The name of the file",
				MaxOccurences: -1,
				// Required:      true,
				Validator: func(arg string) bool {
					return strings.HasSuffix(arg, "/")
				},
			},
			{
				Name:        "file-name",
				Description: "The name of the file",
				Required:    true,
			},
			{
				Name:        "file-extension",
				Description: "The extension of the file if any. Requires a '.' at the start to be valid",
				Validator: func(arg string) bool {
					return strings.HasPrefix(arg, ".")
				},
			},
		},
	},
}

var testArgs = []string{"C:\\path\\to\\file\\betterargs.exe",
	"--files", "./path/to/dir/", "C:/foo/bar/", "filename", ".ext",
	// "--file", "path/to/file.ext", "path2/to3/file45.ext",
	// "--help",
}

func main() {
	fmt.Printf("\n\n---- Start %s ----\n", time.Now().Format(time.TimeOnly))

	fmt.Println(betterargs.ParseArguments(testArgs, argOptions))
}
