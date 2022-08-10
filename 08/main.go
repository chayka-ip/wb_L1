package main

import (
	"reflect"
	"strconv"
)

/*
	Дана переменная int64.
	Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var v int64 = 31 // 11111
	var n uint8 = 2  // bit to set

	// Sets third bit to 0
	Int64SetBit(&v, n, false)
	expected := "11011"
	AssertEqual(expected, int64ToBinaryString(v))

	expected = "11111"
	// Sets third bit to 1
	Int64SetBit(&v, n, true)
	AssertEqual(expected, int64ToBinaryString(v))
}

/*
	Sets n-th bit of v to 1 (true) or 0 (false).
	bitNum is in range [0, 63]. Enumeration goes from left to right.
*/
func Int64SetBit(v *int64, bitNum uint8, bitValue bool) {
	var f int64 = 1 << bitNum
	if bitValue {
		/*
			Bitwise AND
			Sets n-th bit of v to 1 if v or f has this bit set to 1
		*/
		*v |= f
	} else {
		/*
			Bit reset (Bitwise AND OR)
			Sets n-th bit of v to 0 where f has bit set to 1
		*/
		*v &^= f
	}

}

func int64ToBinaryString(v int64) string {
	return strconv.FormatInt(v, 2)
}

func AssertEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic("Assertion failed")
	}
}
