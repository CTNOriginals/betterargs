package betterargs

import (
	"fmt"
	"strings"

	"github.com/CTNOriginals/betterargs/utils"
)

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

func (this Input) displayName() string {
	if this.Required {
		return fmt.Sprintf("<%s>", this.Name)
	}

	return fmt.Sprintf("[%s]", this.Name)
}

func (this Input) Guide() (guide string) {
	guide += fmt.Sprintf("\t%s: max[%d]", this.displayName(), this.MaxOccurences)
	guide += fmt.Sprintf("\n\t  %s", this.Description)
	return guide
}

type InputOrder []Input

func (this InputOrder) String() string {
	var items = map[int]string{}

	for i, item := range this {
		items[i] = item.String()
	}

	return utils.MapToString(items, func(val string) string { return val })
}

func (this InputOrder) displayNames() (names []string) {
	for _, input := range this {
		names = append(names, input.displayName())
	}

	return names
}

func (this InputOrder) Guide() (guide string) {
	var items []string
	for _, input := range this {
		items = append(items, input.Guide())
	}

	return strings.Join(items, "\n")
}
