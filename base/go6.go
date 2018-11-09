package base

import (
	"fmt"
	"time"
)

func produce1(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer1(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int, 2)
	go produce1(ch)
	go consumer1(ch)

	time.Sleep(1 * time.Second)
}
