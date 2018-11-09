package container

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v,len=%d,cap=%d\n",s,len(s),cap(s))
}

func main() {
	var s []int

	for i:=0;i<100;i++{
		//printSlice(s)
		s = append(s,i*2+1)
	}
	fmt.Println(s)


	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16)
	printSlice(s2)

	s3 := make([]int,10,32)
	printSlice(s3)


	//copy 操作
	copy(s2,s1)
	printSlice(s2) //[2 4 6 8 0 0 0 0 0 0]

	//删除 8
	s2 = append(s2[:3],s2[4:]...)
	printSlice(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]


	//删除头部和尾部
	//front := s2[0]
	s2 = s2[1:]
	//println(front)
	printSlice(s2)

	//tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	//println(tail)
	printSlice(s2)





}
