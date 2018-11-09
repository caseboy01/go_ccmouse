package base

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupport operator:" + op) //报错
	}

	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "abc.txt"
	//contents,err := ioutil.ReadFile(filename)
	//if err != nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n",contents)
	//}
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(eval(5, 6, "*"))

	fmt.Println(grade(80))

}
