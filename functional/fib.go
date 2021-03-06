package main

import (
	"bufio"
	"fmt"
	"io"
	"mooc_google_shizhan/functional/fib"
	"strings"
)

//菲波那切数列

//func fibonacci() func()int {
//	a,b := 0,1
//	return func() int {
//		a,b = b,a+b
//		return a
//	}
//}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
	for i := 0; i < 15 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fibonacci()
	//var f intGen = fibonacci()
	var f intGen = fib.Fibonacci()
	printFileContents(f)

}
