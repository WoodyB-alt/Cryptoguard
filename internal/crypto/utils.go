package crypto

import (
	"fmt"
)

// PrintError is a simple helper to format errors
func PrintError(prefix string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", prefix, err)
	}
}
