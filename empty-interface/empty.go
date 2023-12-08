package main

import (
	"fmt"
)

type Number interface {
	int | float64 | uint
}

func main() {

	var v any
	v = 1
	fmt.Printf("v: %v\n", v)

	v = "Hey"
	fmt.Printf("v: %v\n", v)

	// type assertion
	// this is not type conversion
	if value, ok := v.(string); ok {
		fmt.Printf("v is a string: %#v\n", value)
	}

	v = 2023
	// type switch
	switch v.(type) {
	case int:
		fmt.Println("v is an int")
	case string:
		fmt.Println("v is a string")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	fmt.Println("max:", max([]int{1, 2, 3}))
	fmt.Println("max:", max([]float64{1.1, 2.7, 15.3, 0.4, 4.5}))

	fmt.Println("max:", max2([]int{22, 12, 310}))

}

// using generics (with type constraints)
func max[T int | float64](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	mx := nums[0]

	for _, n := range nums[1:] {
		if n > mx {
			mx = n
		}
	}

	return mx
}

// using generics plus interface
func max2[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	mx := nums[0]

	for _, n := range nums[1:] {
		if n > mx {
			mx = n
		}
	}

	return mx
}
