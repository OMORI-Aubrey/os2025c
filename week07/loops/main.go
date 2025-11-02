package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	// i, _ := r.ReadString('\n') // 에러 무시
	i, err := r.ReadString('\n')
	// fmt.Println(err)
	if err != nil {
		log.Fatal(err) // 애러 메시지 보고하고 프로그램 종료
	}
	
	fmt.Println(i)
}
