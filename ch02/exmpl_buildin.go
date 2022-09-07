package main

import (
	"fmt"
)

func main() {
	var (
		s string
		b bool
		i8 int8
		u8 uint8
		y byte
		i16 int16
		u16 uint16
		i32 int32
		r rune
		u32 uint32
		i64 int64
		u64 uint64
		i int
		u uint
		p uintptr
		f32 float32
		f64 float64
		c64 complex64
		c128 complex128
	)

	fmt.Printf(
		`string is a string: %#v
bool is a boolean: %#v
int8 is a 8 bit integer: %#v,
uint8 is a 8 bit unsigned integer: %#v
byte is an alias to uint8: %#v
int16 is a 16 bit integer: %#v
uint16 is a 16 bit unsigned integer: %#v
int32 is a 32 bit integer: %#v
rune is an alias to int32: %#v
uint23 is a 32 bit unsigned integer: %#v
int64 is a 64 bit integer: %#v
uint64 is a 64 bit unsigned integer: %#v
int is a 32 bit integer but not an alias to int32: %#v
uint is a 32 bit unsigned integer but not an alias to uint32: %#v
uintptr is an unsigned integer pointer, type for memory address: %#v
float32 is a 32 bit floating-point number: %#v
float64 is a 64 bit floating-point number: %#v
complex64 is a complex number with 32 bit floating part: %#v
complex128 is a complex number with 64 bit floating part: %#v
`, 
		s, b, i8, u8, y, i16, u16, i32, r, u32, i64, u64, i, u, p, f32, f64, c64, c128,
	)
}
