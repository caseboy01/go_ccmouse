package base

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swapnew(a, b int) (int, int) {
	return b, a
}

func main() {

	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(*pa)
	b := a
	fmt.Println(&b)

	c, d := 3, 4
	swap(&c, &d)
	fmt.Println(c, d)

	e, f := 5, 6
	e, f = swapnew(e, f)
	fmt.Println(e, f)

}
