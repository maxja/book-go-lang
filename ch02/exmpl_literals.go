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
		"with hexadecimal:   \t%E\n",
		withDecimalPoint, withExponent, withHexadec)

	var (
		a        rune = 'a'
		aOctal   rune = '\141'
		aHexadec rune = '\x61'
		a16Hexa  rune = '\u0061'
		a32Hexa  rune = '\U00000061'
	)

	fmt.Printf("runes\n"+
		"character:   \t%c\n"+
		"octal:       \t%c\n"+
		"hexadecimal: \t%c\n"+
		"16-bit hex:  \t%c\n"+
		"32-bit hex:  \t%c\n",
		a, aOctal, aHexadec, a16Hexa, a32Hexa,
	)

	var (
		doubleQuotedString string = "Hello,(\t <tab here>) I'm a string\n"
		rawString          string = `Hello!
I'm a string \t <- it's not a tab anymore.
`
	)

	fmt.Println("strings\n", doubleQuotedString, rawString)
}
