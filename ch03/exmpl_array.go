package main

import "fmt"

func main() {
	var threeElementArray [3]int

	fmt.Printf("threeElementArray declared as `[3]int` and compiler inititlize all elements with zero vlaues %#v\n", threeElementArray)

	var threeElementArrayWithValues = [3]int{10, 20, 30}

	fmt.Printf("threeElementArrayWithValues declared and inititlized with `[3]int{10, 20, 30}`: %#v\n", threeElementArrayWithValues)

	var spacedThreeElementsInArray = [12]int{5: 4, 6, 10: 100}

	fmt.Printf("spacedThreeElementsInArray declared with pointed indexes for some of it's elements `[12]int{5: 4, 6, 10: 100}`: %#v\n", spacedThreeElementsInArray)

	var spreadedArrayElements = [...]int{1, 2, 10: 3}

	fmt.Printf("spreadedArrayelements declared with left off number of elements, replaced with `...`, so compiler will figureout [...]int{1, 2, 10:3}: %#v\n", spreadedArrayElements)

	var (
		leftArray  = [...]int{1, 2, 3}
		rightArray = [3]int{1, 2, 3}
	)

	fmt.Printf("arrays can be compared with `!=` or `==`, [...]int{1,2,3} == [3]int{1,2,3}: %t\n", leftArray == rightArray)

	var twoDemensionalArray [2][3]int = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}

	fmt.Printf("go arrays are one-demensional but they can be nested, like so twoDepensionalArray [2][3]int: %#v\n", twoDemensionalArray)

	fmt.Printf("elements of an array are accessable via their index twoDemesionalArray[1][2]: %d\n", twoDemensionalArray[1][2])
	fmt.Println("elemments counting starts from 0")

	fmt.Printf("number of elements can be retrived via builtin `len` method, len(twoDemensionalArray): %d\n", len(twoDemensionalArray))
	fmt.Println("len would not count elements of nested arrays")
}
