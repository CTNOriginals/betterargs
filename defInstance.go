package betterargs

import (
	"strings"

	"github.com/CTNOriginals/betterargs/utils"
)

type Instance struct {
	// The index at which this instance occurred
	Index int
	// The flag that was used to trigger this instances definition
	Flag string
	// Any arguments that followed that were not defined as flags
	Inputs []string
}

func (this Instance) String() string {
	return utils.StructToString(this)
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
