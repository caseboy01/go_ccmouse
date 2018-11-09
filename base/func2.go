package base

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func apply(op func(int, int) int, a, b int) int {

	//打印函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

}
