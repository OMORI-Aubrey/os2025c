package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	i = strings.TrimSpace(i)               
	num, err := strconv.ParseFloat(i, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func main() {
	fmt.Print("점수 입력: ")

	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	status := ""
	if score >= 90 {
		status = "합격!"
	} else {
		status = "불합격"
	}

	fmt.Printf("%.2f점은 %v\n", score, status)
}