package main

import "fmt"

func main() {
	const (
		a        = 1_000.00
		b string = "peace"
		c int    = 1_000.00
	)

	fmt.Printf("a: %+v\nb: %+v\nc: %+v\n", a, b, c)

	var (
		x float64
	)

	fmt.Println("x variable declared with float64 type: ", x)
	x = a
	fmt.Println("x can be assigned to a const declared with no type: ", x)
	fmt.Println("and can't be assigned to c const declared with type int")
}
