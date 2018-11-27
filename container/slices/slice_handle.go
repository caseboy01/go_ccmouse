package slices

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6] // 2 3 4 5
	s2 := s1[3:5]  //  5 6

	s3 := append(s2, 10) //5 6 10
	s4 := append(s3, 11) // 5 6 10 11
	s5 := append(s4, 12) // 5 6 10 11 12

	fmt.Println(cap(s3), cap(s4), cap(s5))
	fmt.Println(s1, s2, s3, s4, s5)

	fmt.Println(arr) // 0 1 2 3 4 5 6 10

}
