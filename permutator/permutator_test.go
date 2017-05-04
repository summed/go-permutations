package permutator

import (
	"fmt"
	"testing"
)

func TestGetPermutationChannel(t *testing.T) {
	c := GetPermutationChannel([]interface{}{"J", "á", "k", "u", "p"})
	for p := range c {
		fmt.Printf("Chan -> %v\n", p)
	}
}
