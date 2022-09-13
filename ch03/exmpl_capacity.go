package main

import "fmt"

func main() {
	var growSlice []int
	fmt.Printf("Let's declare new array []int, and take a look on it's len: %d, and cap: %d\n", len(growSlice), cap(growSlice))
	fmt.Println("Let's append 5 elements, to see how it's grow")

	for i := 0; i < 5; i++ {
		growSlice = append(growSlice, i)
		fmt.Printf("%d added, and slice looks like: %#v, and len is: %d, cap is: %d\n", i, growSlice, len(growSlice), cap(growSlice))
	}
}
