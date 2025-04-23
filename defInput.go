package betterargs

import "github.com/CTNOriginals/betterargs/utils"

type Input struct {
	Name        string
	Description string

	// Is this input required to be present?
	//
	// If the input is not present after the flag, the program will panic and log an error specifying which input was expected
	Required bool

	// The maximum amount of times that this input can be repeated as long as the validator returns true foe each of them.
	//
	// 0 will be interperted as 1, this is to allow this field to be ommited.
	//
	// -1 for unlimeted
	MaxOccurences int

	// Specify a function that can validate a potential input.
	//
	// If ommited, the input will always be valid given that the arg is not a defined flag
	Validator func(arg string) bool
}

func (this Input) String() string {
	return utils.StructToString(this)
}

type InputOrder []Input

func (this InputOrder) String() string {
	var items = map[int]string{}

	for i, item := range this {
		items[i] = item.String()
	}

	return utils.MapToString(items, func(val string) string { return val })
}

func (this InputOrder) Range() (min int, max int) {
	max = len(this)

	for _, def := range this {
		if def.Required {
			min++
		}
	}

	return min, max
}
