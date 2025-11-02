package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2.34))
	fmt.Println(reflect.TypeOf("Kim Inha"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(19))
}
