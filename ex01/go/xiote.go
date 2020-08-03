// 작성자 : 한진우
// 실행방법 : go run xiote.go
// 주의사항 :

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("What is your name? ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		person := scanner.Text()
		hello := NewHello(person)
		fmt.Println(hello.Message())
		return
	}
}

func NewHello(person string) Hello {
	return Hello{person}
}

type Hello struct {
	Person string
}

func (h *Hello) Message() string {
	return fmt.Sprintf("Hello, %s, nice to meet you!", h.Person)
}
