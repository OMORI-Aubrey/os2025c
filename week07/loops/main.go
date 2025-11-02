package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키를 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수 타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	if score >= 60 {
		fmt.Println(score, "Pass")
	} else {
		fmt.Println(score, "Fail")
	}

	
	// shadowing

	// var float64 float64 = 2.71
	// var f float64 = 3.991
	// fmt.Println(float64)
	// fmt.Println(f)

	// var fmt float64 = 2.71
	// fmt.Println(fmt)
}
