package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	// Declare a nil slice of integers. This slice has no underlying array yet.
	var sl []int
	fmt.Println("length", len(sl))

	// Check if the slice is nil. A nil slice has both length and capacity of 0.
	if sl == nil {
		fmt.Println("nil slice")
	}

	// Declare and initialize a slice with a literal. It has an underlying array.
	sl2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("slice 2: %#v\n", sl2)

	// Create a new slice sl3 by slicing sl2. It includes elements at indices 3 and 4.
	// Note: Slices are zero-indexed and the range is half-open [start, end).
	sl3 := sl2[3:5]
	fmt.Printf("slice 3: %#v\n\n", sl3)

	// Print the length and capacity of slices 'sl2' and 'sl3' in a table format.
	// Length is the number of elements in the slice, and capacity is the total number
	// of elements in the underlying array, starting from the first element of the slice.
	// The capacity of sl3 is 5 because it starts at index 3 of sl2, which has 8 elements.
	// Create a table header
	header := "| Slice | Length | Capacity |"
	divider := "+" + strings.Repeat("-", len(header)-2) + "+"

	// Print the table header and rows
	fmt.Println(divider)
	fmt.Println(header)
	fmt.Println(divider)
	fmt.Printf("| %-5s | %-6d | %-8d |\n", "sl2", len(sl2), cap(sl2))
	fmt.Printf("| %-5s | %-6d | %-8d |\n", "sl3", len(sl3), cap(sl3))
	fmt.Println(divider)
	/* fmt.Println("----------------------------")
	fmt.Println("| Slice | Length | Capacity |")
	fmt.Println("----------------------------")
	fmt.Printf("| sl2   | %-6d | %-8d |\n", len(sl2), cap(sl2))
	fmt.Printf("| sl3   | %-6d | %-8d |\n", len(sl3), cap(sl3))
	fmt.Println("----------------------------")
	*/

	// Print the last 5 entries of the slice sl2 by slicing sl3.
	// The slice sl3[:5] is valid as it is within the capacity of sl3.
	// It extends sl3 to its maximum capacity based on the underlying array (which is sl2).
	fmt.Printf("last 5 entries of sl2 (via sl3) %#v\n", sl3[:5])

	var sl4 []int
	for i := 0; i <= 100; i++ {
		sl4 = appendInt(sl4, i)
	}
	fmt.Printf("sl4: %v\n", sl4)

	var sl5 []string
	for i := 0; i < 5; i++ {
		sl5 = appendString(sl5, "Hello"+fmt.Sprint(i))
	}
	fmt.Printf("\nsl5: %v\n\n", sl5)

	var sl6 []bool
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			sl6 = appendBool(sl6, true)
		} else {
			sl6 = appendBool(sl6, false)
		}
	}
	fmt.Printf("\nsl6: %v\n", sl6)

	fmt.Println(concat([]string{"A", "B", "C"}, []string{"D", "E", "F"}))

	fmt.Println(median([]float64{2, 1, 3, 4, 5}))
	fmt.Printf("TypeOf: %v\n", reflect.TypeOf(2))
	// fmt.Printf("reflect.ArrayOf(2, reflect.TypeOf(2)): %v\n", reflect.ArrayOf(2, reflect.TypeOf(sl6)))

}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("can't get median of an empty slice")
	}
	sort.Float64s(values)
	fmt.Printf("sorted  values: %v\n", values)

	i := len(values) / 2
	fmt.Printf("i: %v\n", i)
	fmt.Printf("(len(values) mod 2): %v\n", len(values)%2)

	if len(values)%2 == 1 {
		return values[i], nil
	}
	fmt.Printf("values[i-1]: %v\n", values[i-1])
	fmt.Printf("values[i]: %v\n", values[i])

	v := (values[i-1] + values[i]) / 2

	return v, nil
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}

// appendInt appends an integer 'v' to the slice 's'.
func appendInt(s []int, v int) []int {
	// Get the index we want to append to, initially set to the length of 's'.
	index := len(s) // 0 indexed

	// Check if there is enough capacity in the slice to append the new element.
	if len(s) < cap(s) {
		// If there's room in the slice's capacity, extend the slice by one element.
		s = s[:len(s)+1]
	} else {
		// If there's not enough capacity, reallocate and copy the slice.
		fmt.Printf("\nreallocate: %d->%d\n", len(s), 2*len(s)+1)

		// Create a new slice 's2' with doubled capacity and one more element.
		s2 := make([]int, 2*len(s)+1)

		// Copy the elements from the old slice 's' to the new slice 's2'.
		copy(s2, s)

		// Update the reference of 's' to the new slice 's2' and extend it by one element.
		s = s2[:len(s)+1]
	}

	// Set the value 'v' at the computed index in the slice.
	s[index] = v

	// Return the modified slice.
	return s
}

func appendString(s []string, v string) []string {
	i := len(s)

	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		s2 := make([]string, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}

func appendBool(s []bool, v bool) []bool {
	i := len(s)
	fmt.Printf("cap 1 of bool slice in %v iteration: %v\n", i, cap(s))

	if cap(s) > len(s) {
		s = s[:len(s)+1]
	} else {
		s2 := make([]bool, 2*len(s)+1)

		copy(s2, s)

		s = s2[:len(s)+1]

		fmt.Printf("cap 2 of bool slice in %v iteration : %v\n", i, cap(s))
	}

	s[i] = v

	return s
}
