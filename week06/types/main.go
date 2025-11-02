package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 64f float64 // Error
	var f64 float64

	// total_price := 1000
	// totalPrice := 1000
	totalPrice := 1000

	fmt.Println(totalPrice)
	fmt.Println(f64, reflect.TypeOf(f64))

}
