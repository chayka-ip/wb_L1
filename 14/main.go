package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

/*
	Разработать программу, которая в рантайме способна определить тип
	переменной: int, string, bool, channel из переменной типа interface{}.
*/

type mint int

func main() {

	var i interface{}
	fmt.Println(ObtainType(mint(1)))             // Output: unknown type
	fmt.Println(ObtainType(new(int)))            // Output: unknown type
	fmt.Println(ObtainType(new(chan int)))       // Output: unknown type
	fmt.Println(ObtainType(nil))                 // Output: unknown type
	fmt.Println(ObtainType(i))                   // Output: unknown type
	fmt.Println(ObtainType(time.Now()))          // Output: unknown type
	fmt.Println(ObtainType(1))                   // Output: int
	fmt.Println(ObtainType(false))               // Output: bool
	fmt.Println(ObtainType("abc"))               // Output: string
	fmt.Println(ObtainType(make(chan int)))      // Output: chan int
	fmt.Println(ObtainType(make(chan string)))   // Output: chan string
	fmt.Println(ObtainType(make(chan struct{}))) // Output: chan struct {}

}

func ObtainType(v interface{}) string {
	if pt, err := obtainTypeWithReflect(v); err == nil {
		return pt
	}

	if ct, err := obtainChannelType(v); err == nil {
		return ct
	}
	return "unknown type"
}

func obtainTypeWithReflect(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t != nil {
		switch t {
		case reflect.TypeOf(1), reflect.TypeOf(true), reflect.TypeOf(""):
			return t.Name(), nil
		}
	}
	return "", errors.New("unknown primitive type")
}

// alternative way to check for type
func obtainTypeWithAssert(v interface{}) (string, error) {
	switch v.(type) {
	case int:
		return "int", nil
	case bool:
		return "bool", nil
	case string:
		return "string", nil
	}
	return "", errors.New("unknown primitive type")
}

func obtainChannelType(v interface{}) (string, error) {
	t := reflect.TypeOf(v)
	if t != nil {
		// check for channel
		isChan := t.Kind() == reflect.Chan
		if isChan {
			// retrieve channel type
			chanType := reflect.ValueOf(v).Type().Elem()
			return fmt.Sprint("chan ", chanType), nil
		}
	}
	return "", errors.New("unknown channel")
}
