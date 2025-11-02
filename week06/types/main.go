package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Zero Values
	var f64 float64
	var t bool
	var s string
	var i int
	var i16 int16

	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(i16, reflect.TypeOf(i16))
}