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

		var instance = Instance{
			Index: i,
			Flag:  arg,
		}

		if parsed.Args[key] == nil {
			parsed.Args[key] = make([]Instance, 0)
		}

		parsed.Args[key] = append(parsed.Args[key], instance)

		var minInputs, maxInputs = def.Inputs.Range()
		if minInputs == 0 && maxInputs == 0 {
			continue
		}
	}

	return parsed
}
