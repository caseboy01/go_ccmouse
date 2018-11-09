package base

import "fmt"

//func div(a,b int)(int,int) {
//	return a/b,a%b
//}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	q, r := div(13, 4)
	fmt.Println(q, r)

	fmt.Println(sum(1, 2, 3, 4, 5))

}
