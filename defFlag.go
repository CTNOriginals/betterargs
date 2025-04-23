package betterargs

import (
	"fmt"
	"slices"

	"github.com/CTNOriginals/betterargs/utils"
)

type Flag struct {
	// All possible flags that can trigger this option
	//  // for --help
	//  []string{"-H", "-?"}
	Aliases []string

	Description string

	// The amount of input arguments after the flag to expect.
	Inputs InputOrder
}

func (this Flag) String() string {
	return utils.StructToString(this)
}

// A map of argument options that define all options for the current program.
// The key of the map will also be its default usage as an argument but with '--' infront of it
//
//	{ "flag-name": MFlags{...} }
type MFlags map[string]Flag

func (this MFlags) String() string {
	return utils.MapToString(this, func(val Flag) string { return val.String() })
}

func (this MFlags) Find(flag string) (name string, definition Flag, err error) {
	for key, def := range this {
		if flag == "--"+key || slices.Contains(def.Aliases, flag) {
			return key, def, nil
		}
	}

	return "", Flag{}, fmt.Errorf("Argument flag '%s' is not defined", flag)
}
