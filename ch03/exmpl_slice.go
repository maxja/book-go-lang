package main

import "fmt"

func main() {
	var sliceButNotArray = []int{1, 2, 3}
	fmt.Printf("slice can be declared and inititalized almost exactly like array, except size should be omited []int{1,2,3}: %#v\n", sliceButNotArray)

	var spacedSlice = []int{5: 4, 6, 10: 100}
	fmt.Printf("slice can be intitlized spaced by declaring indexes for values, []int{5:4, 6, 10:100}: %#v\n", spacedSlice)

	var multidimensionalSlice [][]int
	fmt.Printf("slice can be nested with other slices [][]int: %#v\n", multidimensionalSlice)

	var x = []int{0}
	x[0] = 10
	fmt.Printf("slice x = []int{0} elements can be assigned referencing by index x[0] = 10: %#v\n", x)

	var nilSlice []int
	fmt.Printf("by declaring without initializing nil slice will be created: %#v\n", nilSlice)

	var leftSlice = []int{1, 2, 3}
	fmt.Printf("slices can't be compare one to another, they are only comparable to `nil` []int{1,2,3} != nil: %t, []int == nil: %t\n", leftSlice != nil, nilSlice == nil)

	fmt.Printf("len method works for slice too, len([]int{1, 2, 3}): %d\n", len(leftSlice))
	fmt.Printf("nil slice will return 0 len(nilSlice): %d\n", len(nilSlice))

	fmt.Printf("slice can be appended with element of same type var x = []int{10}: %#v\n", x)
	x = append(x, 20)
	fmt.Printf("builtin function `append` can be used for that purpose, result of should be assigned to the same variable, x = append(x, 20): %#v\n", x)
	x = append(x, 30, 40, 50)
	fmt.Printf("append function accepts slice as 1st argument, and one or more elements to add, x = append(x, 30, 40, 50): %#v\n", x)

	y := []int{100, 200, 300}
	x = append(x, y...)
	fmt.Printf("go have variadic operator, that can help to combine two or more slices together, y := []int{100, 200, 300}; x = append(x, y...): %#v\n", x)
}
