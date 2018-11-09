package base

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Print(a,s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "ghi"
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b)) //不需要强制转换成float了
	fmt.Println(filename, c)
}

func enums() {
	//const(
	//	java = 0
	//	php = 1
	//	golang = 2
	//	cpp = 3
	//)
	const (
		java = iota
		php
		golang
		cpp
	)

	//iota 自增值类型
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(java, php, golang, cpp)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()
	triangle()
	consts()
	enums()
}
