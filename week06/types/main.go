package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var id int16
	// var name string
	// var gpa float32

	// id = 999
	// name = "Kim Inha"
	// gpa = 3.99

	// var id int16 = 999
	// var name string = "Kim Inha"
	// var gpa float32 = 3.99

	// var id = 999
	// var name = "Kim Inha"
	// var gpa = 3.99

	id := 999
	name := "Kim Inha"
	gpa := 3.99

	fmt.Print("학번은", id, reflect.TypeOf(id))
	fmt.Println("이름은", name, reflect.TypeOf(name))
	fmt.Println("평점:", gpa, reflect.TypeOf(gpa))

}
