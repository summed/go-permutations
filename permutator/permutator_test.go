package permutator

import (
	"fmt"
	"log"
	"testing"
)

// TODO
func TestGetPermutationChannel(t *testing.T) {
	d, err := toInterfaceSlice([]byte("Demmus"))
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	c := GetPermutationChannel(d)
	for p := range c {
		fmt.Printf("Chan -> %#v - ", p)
		for i := 0; i < len(p); i++ {
			fmt.Printf("%s", string(p[i].(byte)))
		}
		fmt.Printf("\n")
	}
}
