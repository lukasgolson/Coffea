package application

import (
	"fmt"
	"strings"
)

func FormatNumbers(intArray []int, itemsPerLine int) string {
	var formattedNumbers []string

	for i, b := range intArray {
		// Append the decimal number with zero padding to the hundreds place
		formattedNumbers = append(formattedNumbers, fmt.Sprintf("%03d", b))

		// Check if we need to add a newline
		if (i+1)%itemsPerLine == 0 && i != len(intArray)-1 {
			formattedNumbers = append(formattedNumbers, "\n")
		} else if i != len(intArray)-1 {
			// Add a comma if it's not the last item
			formattedNumbers = append(formattedNumbers, " ")
		}
	}

	return strings.Join(formattedNumbers, "")
}
