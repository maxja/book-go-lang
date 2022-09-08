package main

import "fmt"

func main() {
	var (
		i int     = 1_123
		u uint16  = 3_000
		f float64 = 3.14
	)

	fmt.Println("Explicit Type Conversion")
	fmt.Printf("Let say, we have int `i`=%d, uint `u`=%d, and float `f`=%f\n", i, u, f)
	fmt.Printf("To apply math operation like this: f *= u - i\n")
	f *= float64(int(u) - i)
	fmt.Printf("We have to convert `u` and the result of subtraction: f *= float64(int(u) - i) = %f\n", f)
}
