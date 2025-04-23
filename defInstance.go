package betterargs

import (
	"strings"

	"github.com/CTNOriginals/betterargs/utils"
)

type MInputInstances map[string][]string

func (this MInputInstances) String() string {
	return utils.MapToString(this, func(val []string) string {
		return strings.Join(val, "\n")
	})
}

type Instance struct {
	// The index at which this instance occurred
	Index int
	// The arg that was used to trigger this instances definition
	Arg string
	// All the args that followed that are defined and validated to be inputs for this flag
	Inputs MInputInstances
}

func (this Instance) String() string {
	return utils.StructToString(this)
}

func (this *Instance) PushInput(name string, value string) {
	this.Inputs[name] = append(this.Inputs[name], value)
}

// All arg instances grouped by their arg definition key
type MInstances map[string][]Instance

func (this MInstances) String() string {
	return utils.MapToString(this, func(val []Instance) string {
		var stringified = make([]string, len(val))
		for i, item := range val {
			stringified[i] = item.String()
		}
		return strings.Join(stringified, ",\n\n")
	})
}

func (this MInstances) NewInstance(key string, index int, arg string) (instance *Instance) {
	instance = &Instance{
		Index:  index,
		Arg:    arg,
		Inputs: map[string][]string{},
	}

	//? Add the new instance to the list
	this[key] = append(this[key], *instance)
	//? Reteive the instance from the array so that any future changes to it are updated in the array automatically
	instance = &this[key][len(this[key])-1]

	return instance
}
