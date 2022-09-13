package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Printf("New slice can be initialized with given length make([]int, 5): %#v\n", x)
	fmt.Printf("capacity (%d) will be equal to length (%d)\n\n", cap(x), len(x))

	x = append(x, 10)
	fmt.Printf("Appending to this slice x = append(x, 10) will add new element at the end, and enlarge length (%d) and capacity (%d): %#v\n\n", len(x), cap(x), x)

	x[0] = 1
	fmt.Printf("To fill in preallocated elements, indexing should be used x[0] = 1: %#v\n\n", x)

	y := make([]int, 0, 10)
	fmt.Printf("Append can be used to fill entire slice, but length should be set to 0, and capacity give make([]int, 0, 10): %#v\n", y)
	y = append(y, 1, 2, 3)
	fmt.Printf("so that new elements will be putted on empty places y = append(y, 1, 2, 3): %#v\n", y)
}
