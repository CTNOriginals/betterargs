package betterargs

import (
	"fmt"
	"slices"

	"github.com/CTNOriginals/betterargs/utils"
)

type Flag struct {
	Description string

	// All possible flags that can trigger this option
	//  // for --help
	//  []string{"-H", "-?"}
	Aliases []string

	// All the inputs in the order to expect them in
	//
	// If, for example, the first input's validator returned false
	// and the input after it returned true,
	// the next arg will be evaluated by the third input
	Inputs InputOrder

	Action func(instance Instance)
}

func (this Flag) String() string {
	return utils.StructToString(this)
}

// A map of argument options that define all options for the current program
//
// The key of the map will also be its default usage as an argument but with '--' infront of it
//
//	{ "flag-name": MFlags{...} }
type MFlags map[string]Flag

func (this MFlags) String() string {
	return utils.MapToString(this, func(val Flag) string { return val.String() })
}

func (this MFlags) validate() {
	for key, def := range this {
		for i, input := range def.Inputs {
			if input.MaxOccurences == 0 {
				input.MaxOccurences = 1
			}

			def.Inputs[i] = input
		}

		this[key] = def
	}
}

func (this MFlags) find(flag string) (name string, definition Flag, err error) {
	for key, def := range this {
		if flag == "--"+key || slices.Contains(def.Aliases, flag) {
			return key, def, nil
		}
	}

	return "", Flag{}, fmt.Errorf("Argument flag '%s' is not defined", flag)
}

// Parse all valid inputs into their respective input definition
//
// args should be the remaining arguments after the flag that triggered this definition
//
//	var args = ["script/path/file.go", "--file", "dir/", "file.ext"]
//	definition.parseInputs(args[2:]) // args = ["dir/", "file.ext"]
//
// This function returns the amount of inputs that were successfully parsed from args.
func (this MFlags) parseInputs(def Flag, instance *Instance, args []string) (offset int) {
	//TODO print the help section for this flag along with any panic

	for i := 0; i < len(def.Inputs); i++ {
		var input = def.Inputs[i]

		if offset >= len(args) {
			break
		}

		var arg = args[offset]
		_, _, err := this.find(arg)
		var argIsFlag = err == nil

		if (input.Validator != nil && !input.Validator(arg)) || (input.Validator == nil && argIsFlag) {
			if input.Required && len(instance.Inputs[input.Name]) == 0 {
				panic(fmt.Sprintf("Argument Input ERROR: Expected valid input argument for '%s' but found invalid argument: '%s'", input.Name, arg))
			}
			continue
		}

		instance.pushInput(input.Name, arg)
		offset++

		if input.MaxOccurences == -1 || len(instance.Inputs[input.Name]) < input.MaxOccurences {
			i--
		}
	}

	for _, input := range def.Inputs {
		if input.Required && len(instance.Inputs[input.Name]) == 0 {
			panic(fmt.Sprintf("Argument Input ERROR: Expected valid input argument for '%s'", input.Name))
		}
	}

	return offset
}
