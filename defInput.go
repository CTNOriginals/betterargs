package betterargs

import "github.com/CTNOriginals/betterargs/utils"

type Input struct {
	Description string
	Required    bool
	Validator   func(arg string) bool
}

func (this Input) String() string {
	return utils.StructToString(this)
}

type MInputs map[string]Input

func (this MInputs) String() string {
	return utils.MapToString(this, func(val Input) string { return val.String() })
}

func (this MInputs) Range() (min int, max int) {
	max = len(this)

	for _, def := range this {
		if def.Required {
			min++
		}
	}

	return min, max
}
