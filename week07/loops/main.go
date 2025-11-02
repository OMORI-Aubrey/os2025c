package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n') // 에러 무시
	fmt.Println(i)
}
