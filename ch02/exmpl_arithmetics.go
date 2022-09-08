package main

import (
	"fmt"
)

func main() {
	const (
		x, y int = 10, 3
	)

	fmt.Println("Arithmetic operations")
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x %% y = %d\n", x%y)

	var z int = 5

	fmt.Println("Atomic arithmetic operations")
	fmt.Printf("z = %d, y = %d\n", z, y)
	z += y
	fmt.Printf("z += y, z = %d\n", z)
	z -= y
	fmt.Printf("z -= y, z = %d\n", z)
	z *= y
	fmt.Printf("z *= y, z = %d\n", z)
	z /= y
	fmt.Printf("z /= y, z = %d\n", z)
	z %= y
	fmt.Printf("z %%= y, z = %d\n", z)

	fmt.Println("Comparison")
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("x == y: %d\n", x == y)
	fmt.Printf("x != y = %d\n", x != y)
	fmt.Printf("x >= y = %d\n", x >= y)
	fmt.Printf("x <= y = %d\n", x <= y)

	z = 11
	fmt.Println("Bitwise operations")
	fmt.Printf("z = %b, y = %b\n", z, y)
	fmt.Printf("z & y, z = %b\n", z & y)
	fmt.Printf("z | y, z = %b\n", z | y)
	fmt.Printf("z ^ y, z = %b\n", z ^ y)
	fmt.Printf("z &^ y, z = %b\n", z &^ y)
	fmt.Printf("z << y, z = %b\n", z << y)
	fmt.Printf("z >> y, z = %b\n", z >> y)
}
