package betterargs

import (
	"fmt"

	"github.com/CTNOriginals/betterargs/utils"
)

type ParsedArguments struct {
	//TODO Find better wording for this comment:
	// The unedited arguments as they were passed in at the launch of the program
	Raw []string

	// The path to the file that is currently running.
	SourceFile string

	// The definitions of all possible argument flags and their implimentations
	Definitions MFlags

	// All arg instances that were present
	Args MInstances
}

func (this ParsedArguments) String() string {
	return utils.StructToString(this)
}

func ParseArguments(args []string, definitions MFlags) (parsed ParsedArguments) {
	definitions.Validate()

	parsed.Raw = make([]string, len(args))
	copy(parsed.Raw, args)

	parsed.SourceFile = args[0]
	args, _ = utils.Splice(args, 0, 1)

	parsed.Definitions = definitions
	parsed.Args = MInstances{}

	for i := 0; i < len(args); i++ {
		var arg = args[i]
		var key, def, err = definitions.Find(arg)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var instance = &Instance{
			Index:  i,
			Flag:   arg,
			Inputs: map[string][]string{},
		}

		//? Add the new instance to the list
		parsed.Args[key] = append(parsed.Args[key], *instance)
		//? Reteive the instance from the array so that any changes to it are also present in the array automatically
		instance = &parsed.Args[key][len(parsed.Args[key])-1]

		if parsed.Args[key] == nil {
			parsed.Args[key] = make([]Instance, 0)
		}

		var minInputs, maxInputs = def.Inputs.Range()
		if minInputs == 0 && maxInputs == 0 {
			continue
		}

		//? The amout of indexes to look ahead for inputs
		//? This amount goes up each time an input is succesfully validated by any input definition
		var offset = 1
		for j := 0; j < len(def.Inputs); j++ {
			var input = def.Inputs[j]
			var nextIndex = i + offset

			if nextIndex >= len(args) {
				break
			}

			var nextArg = args[nextIndex]

			//?? I dont know how to combine the two conditions without making it scream at me :/
			if input.Validator != nil && !input.Validator(nextArg) {
				continue
			} else if _, _, err := definitions.Find(nextArg); err == nil {
				continue
			}

			instance.PushInput(input.Name, nextArg)
			offset++

			if input.MaxOccurences > -1 && len(instance.Inputs[input.Name]) < input.MaxOccurences {
				j--
			}
		}

		i += offset - 1
	}

	fmt.Println()

	return parsed
}
