package betterargs

import (
	"fmt"
	"slices"

	"github.com/CTNOriginals/betterargs/utils"
)

type Flags []string

type Arg struct {
	// All possible flags that can trigger this option
	//  betterargs.Flags{"--help", "-H", "?"}
	Flags Flags

	Description string

	// The amount of input arguments after the flag to expect.
	Inputs MInputs

	// // Should this argument be exclusively present?
	// // If true while it is not the first argument flag,
	// // Any arguments that follow a
	// Exclusive bool
}

func (this Arg) String() string {
	return utils.StructToString(this)
}

// A map of argument options that define all options for the current program
//
//	{ "option-name": MArgs{...} }
type MArgs map[string]Arg

func (this MArgs) String() string {
	return utils.MapToString(this, func(val Arg) string { return val.String() })
}

// = Consept
func (this MArgs) Push(name string, def Arg) {
	this[name] = def

	//* Execute arg parser again with the new def
	//> This has a weird application, if, throughout the program an arg would be defined, a lot of args would be redifined a lot of the time...

	/*
		logVerbose = defs.push(--verbose)
		if lobVerbose { print("some verpose data") }
		applyData()

		func applyData() {
			logVerbose = defs.push(--verbose) //! second definition

			[Arg(--verbose)]
			print("data applied!")
		}
	*/
}

func (this MArgs) Find(flag string) (name string, definition Arg, err error) {
	for key, def := range this {
		if slices.Contains(def.Flags, flag) {
			return key, def, nil
		}
	}

	return "", Arg{}, fmt.Errorf("Argument flag '%s' is not defined", flag)
}
