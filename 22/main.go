package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две
	числовых переменных a,b, значение которых > 2^20.
*/

var (
	MIN_NUMBER_LEN = 50
	MAX_NUMBER_LEN = 100
)

func main() {
	/*
		math/big allows to perform computations
		with numbers of arbitrary length and precision
	*/

	a, b := getRandomBigInt(), getRandomBigInt()
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	r := big.NewInt(0)
	fmt.Println("Mul: ", r.Mul(a, b))
	fmt.Println("Div: ", r.Div(a, b))
	fmt.Println("Sum: ", r.Add(a, b))
	fmt.Println("Sub: ", r.Sub(a, b))
}

func getRandomBigInt() *big.Int {
	s := getNumericString(MIN_NUMBER_LEN, MAX_NUMBER_LEN)
	//convert numeric string to big number
	n, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		log.Fatal("Unable to convert string to number: ", s)
	}
	return n
}

func getNumericString(minLength, maxLength int) string {
	// set new seed to generate reload randomness
	rand.Seed(time.Now().UnixNano())
	// generate string length
	n := rand.Intn(maxLength-minLength+1) + minLength

	var b strings.Builder
	// required size is known, do allocation once
	b.Grow(n)

	// generate numeric string by adding digits one by one
	for i := 0; i < n; i++ {
		v := fmt.Sprint(rand.Intn(10))
		b.WriteString(v)
	}
	return b.String()
}
