package formatting

import (
	"fmt"
	"strings"
)

type Thing struct {
	Name string
	Version string
	Description string
}

func FixedColumnOutput(data []Thing) (output []string) {
	output = make([]string, 0)

	largest := 0
	spacing := 4
	pad := strings.Repeat(" ", spacing)

	for _, thing := range data {
		if len(thing.Name) > largest {
			largest = len(thing.Name)
		}
	}

	fieldLength := largest + spacing

	for _, thing := range data {

		// %% is literal %

		// Have to construct the format string first
		formatstring := fmt.Sprintf("%%%ds%%s%%s%%s%%s\n", fieldLength)

		// Then can fill with variables
		output = append(output, fmt.Sprintf(formatstring, thing.Name, pad, thing.Version, pad, thing.Description ))
	}

	return output
}