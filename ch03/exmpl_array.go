package main

import "fmt"

func main() {
	var threeElementArray [3]int
	fmt.Printf("threeElementArray declared as `[3]int` and compiler initialize all elements with zero values %#v\n\n", threeElementArray)

	var threeElementArrayWithValues = [3]int{10, 20, 30}
	fmt.Printf("threeElementArrayWithValues declared and initialized with `[3]int{10, 20, 30}`: %#v\n\n", threeElementArrayWithValues)

	var spacedThreeElementsInArray = [12]int{5: 4, 6, 10: 100}
	fmt.Printf("spacedThreeElementsInArray declared with pointed indexes for some of it's elements `[12]int{5: 4, 6, 10: 100}`: %#v\n\n", spacedThreeElementsInArray)

	var spreadedArrayElements = [...]int{1, 2, 10: 3}
	fmt.Printf("spreadedArrayelements declared with left off number of elements, replaced with `...`, so compiler will calculate [...]int{1, 2, 10:3}: %#v\n\n", spreadedArrayElements)

	var (
		leftArray  = [...]int{1, 2, 3}
		rightArray = [3]int{1, 2, 3}
	)
	fmt.Printf("arrays can be compared with `!=` or `==`, [...]int{1,2,3} == [3]int{1,2,3}: %t\n\n", leftArray == rightArray)

	var twoDemensionalArray [2][3]int = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Printf("go arrays are one-dimensional but they can be nested, like so twoDepensionalArray [2][3]int: %#v\n\n", twoDemensionalArray)

	fmt.Printf("elements of an array are accessible via their index twoDemesionalArray[1][2]: %d\n", twoDemensionalArray[1][2])
	fmt.Println("elements counting starts from 0")
	twoDemensionalArray[1][2] = 100
	fmt.Printf("and existing element can be re-assigned with value of the same type `twoDemensionalArray[1][2] = 100`: %d\n\n", twoDemensionalArray[1][2])
	fmt.Printf("number of elements can be retrieved via builtin `len` method, `len(twoDemensionalArray)`: %d\n\n", len(twoDemensionalArray))
	fmt.Println("len would not count elements of nested arrays")
}
