package permutator

// factorial returns factorial of n
func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

// getCopyOfSliceWithoutElemAt returns a slice continating the input elements of a, except the one at index i
func getCopyOfSliceWithoutElemAt(a []interface{}, i int) []interface{} {
	r := make([]interface{}, len(a)-1)
	ptr := 0
	for p, d := range a {
		if p == i {
			continue
		}
		r[ptr] = d
		ptr++
	}
	return r
}

// GetPermutationChannel returns a []interface{} channel containing permutations of the input slice
func GetPermutationChannel(s []interface{}) <-chan []interface{} {
	rc := make(chan []interface{})
	go func() {
		defer close(rc)
		var n, l int

		l = len(s) // Size of each permutation set
		n = l - 1  // Size of recursion subset (minus a currently processing element)

		if n > 0 {
			for i := 0; i < l; i++ { // Recursively get permutation
				t := getCopyOfSliceWithoutElemAt(s, i) // subset slice of interfaces (minus the currently processing element)
				g := GetPermutationChannel(t)
				for x := range g {
					var r = make([]interface{}, l) // return (channeled) slice of interfaces
					r[0] = s[i]                    // Insert the 'skipped' element as first element in return slice
					ptr := 1
					for y := 0; y < len(x); y++ {
						r[ptr] = x[y]
						ptr++
					}
					rc <- r
				}
			}

		} else {
			rc <- s
		}
	}()
	return rc
}
