package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month() // month := now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	// var univ string = "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	// changed := changer.Replace(univ)
	changed := changer.Replace("Go$ Inha$")
	fmt.Println(changed)

}
