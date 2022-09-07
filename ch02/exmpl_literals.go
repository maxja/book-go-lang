package main

import "fmt"

func main() {
	var (
		ten     uint8 = 10
		binary  uint8 = 0b10
		octal   uint8 = 0o10
		hexadec uint8 = 0x10
	)

	fmt.Printf("integers\n"+
		"base:   \tvalue \tin base ten\n"+
		"ten:    \t%d \t%d\n"+
		"binary: \t%b \t%d\n"+
		"octal:  \t%o \t%d\n"+
		"hexadecimal: \t%x \t%d\n",
		ten, ten, binary, binary, octal, octal, hexadec, hexadec)

	var bigNumber uint16 = 65_535

	fmt.Println("big number literal written with _ between thousands: ", bigNumber)

	var (
		withDecimalPoint float32 = 0.123_321
		withExponent     float32 = -1.23e-1
		withHexadec      float64 = 0x0.0FCp1
	)

	fmt.Printf("floats\n"+
		"with decimal point: \t%f\n"+
		"with exponent set:  \t%e\n"+
		"with hexadecimal:   \t%E",
		withDecimalPoint, withExponent, withHexadec)
}
