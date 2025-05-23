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
	definitions.validate()

	parsed.Raw = make([]string, len(args))
	copy(parsed.Raw, args)

	parsed.SourceFile = args[0]
	args, _ = utils.Splice(args, 0, 1)

	parsed.Definitions = definitions
	parsed.Args = MInstances{}

	for i := 0; i < len(args); i++ {
		var arg = args[i]
		var key, def, err = definitions.find(arg)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var instance = parsed.Args.newInstance(key, i, arg)

		i += definitions.parseInputs(def, instance, args[i+1:])
	}

	return parsed
}

func ExecuteArguments(parsedArgs ParsedArguments) {
	for key, instances := range parsedArgs.Args {
		var def = parsedArgs.Definitions[key]

		if def.Action == nil {
			continue
		}

		for _, instance := range instances {
			def.Action(instance)
		}
	}
}
