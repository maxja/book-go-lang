package main

import "fmt"

func main() {
	const (
		pupa string = "pupa"
		lupa string = "lupa"
	)

	fmt.Printf("string comparison: %s and %s\n", pupa, lupa)
	fmt.Printf("is %s == %s: %t\n", pupa, lupa, pupa == lupa)
	fmt.Printf("is %s != %s: %t\n", pupa, lupa, pupa != lupa)
	fmt.Printf("is %s <= %s: %t\n", pupa, lupa, pupa <= lupa)
	fmt.Printf("is %s >= %s: %t\n", pupa, lupa, pupa >= lupa)
}
