package main

import (
	"bufio"
	"fmt"
	"mooc_google_shizhan/functional/fib"
	"os"
)

// 先进后出
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func tryDeferd() {

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("ok..............")
		}
	}

}

func writeFile(filename string) {

	//file,err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("this is test error") //自定义error
	if err != nil {
		//panic(err)
		//fmt.Println(err.Error())
		fmt.Println(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err) //open fib.txt The file exists.
		}

		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	//tryDefer()
	//tryDeferd()
	writeFile("fib.txt")

}
