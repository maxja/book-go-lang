package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		s    string
		b    bool
		i8   int8
		u8   uint8
		y    byte
		i16  int16
		u16  uint16
		i32  int32
		r    rune
		u32  uint32
		i64  int64
		u64  uint64
		i    int
		u    uint
		p    uintptr
		f32  float32
		f64  float64
		c64  complex64
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
int is a CPU dependent integer: %#v
uint is a CPU dependent unsigned integer: %#v
uintptr is an unsigned integer pointer, type for memory address: %#v
float32 is a 32 bit floating-point number: %#v
float64 is a 64 bit floating-point number: %#v
complex64 is a complex number with 32 bit floating part: %#v
complex128 is a complex number with 64 bit floating part: %#v
`,
		s, b, i8, u8, y, i16, u16, i32, r, u32, i64, u64, i, u, p, f32, f64, c64, c128,
	)

	var (
		uninitializedBoolFlag bool
		initializedBoolFlag   bool = true
	)

	fmt.Println("uninitialized boolean variable will get: ", uninitializedBoolFlag)
	fmt.Println("initialized get what you assign, true in this case: ", initializedBoolFlag)

	const (
		maxUint   uint   = ^uint(0)
		minUint   uint   = 0
		maxInt    int    = int(maxUint >> 1)
		minInt    int    = -maxInt - 1
		maxUint8  uint8  = ^uint8(0)
		minUint8  uint8  = 0
		maxInt8   int8   = int8(maxUint8 >> 1)
		minInt8   int8   = -maxInt8 - 1
		maxUint16 uint16 = ^uint16(0)
		minUint16 uint16 = 0
		maxInt16  int16  = int16(maxUint16 >> 1)
		minInt16  int16  = -maxInt16 - 1
		maxUint32 uint32 = ^uint32(0)
		minUint32 uint32 = 0
		maxInt32  int32  = int32(maxUint32 >> 1)
		minInt32  int32  = -maxInt32 - 1
		maxUint64 uint64 = ^uint64(0)
		minUint64 uint64 = 0
		maxInt64  int64  = int64(maxUint64 >> 1)
		minInt64  int64  = -maxInt64 - 1
	)

	fmt.Printf("%-13s %-32s | %-32s\n", "type", "min", "max")
	fmt.Printf("%-12s: %-32d | %-32d\n", "uint", minUint, maxUint)
	fmt.Printf("%-12s: %-32d | %-32d\n", "int", minInt, maxInt)
	fmt.Printf("%-12s: %-32d | %-32d\n", "uint8", minUint8, maxUint8)
	fmt.Printf("%-12s: %-32d | %-32d\n", "int8", minInt8, maxInt8)
	fmt.Printf("%-12s: %-32d | %-32d\n", "uint16", minUint16, maxUint16)
	fmt.Printf("%-12s: %-32d | %-32d\n", "int16", minInt16, maxInt16)
	fmt.Printf("%-12s: %-32d | %-32d\n", "uint32", minUint32, maxUint32)
	fmt.Printf("%-12s: %-32d | %-32d\n", "int32", minInt32, maxInt32)
	fmt.Printf("%-12s: %-32d | %-32d\n", "uint64", minUint64, maxUint64)
	fmt.Printf("%-12s: %-32d | %-32d\n", "int64", minInt64, maxInt64)

	fmt.Printf("%-12s: %-.26e | %-.26e\n", "float32", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("%-12s: %-.25e | %-.25e\n", "float64", math.SmallestNonzeroFloat64, math.MaxFloat64)
}
