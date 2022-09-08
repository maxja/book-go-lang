package main

import "fmt"

var outOfFunctionScope bool = true
var isInitialized bool

func main() {
	inFunctionScope := true

	var a, b float64 = 0, 0

	var (
		x = 0
		y float64
		z float64 = 0
	)

	i, j := 'i', 'j'

	fmt.Println("Variables declaration and initialization")
	fmt.Println("outOfFunctionScope variable was declared with `var` keyword and assigned with true: ", outOfFunctionScope)
	fmt.Println("isInitialized variable was declared with `var` keyword but not assigned: ", isInitialized)
	fmt.Println("inFunctionScope variable was declared and assigned via `:=` and true: ", inFunctionScope)
	fmt.Println("a, b variables was declared via `var` with type float and assigned to 0, 0: ", a, ", ", b)
	fmt.Println("x, y, z variables was declared via `var` scope, x without type but assigned to 0, y with type float64 but without assigned value, and z with float64 and assigned to value 0: ", x, ", ", y, ", ", z)
	fmt.Println("i, j variables was declared and assigned via `:=` with values 'i', 'j': ", i, ", ", j)
}
