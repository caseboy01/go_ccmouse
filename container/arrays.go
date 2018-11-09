package container

import "fmt"

func printArray(arr *[5]int)  {
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}



func main() {

	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int //4行5列

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//遍历数组1
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}
	//遍历数组2
	for i:=range arr3{
		fmt.Println(arr3[i])
	}

	//遍历数组3
	for _,v :=range arr3{
		fmt.Println(v)
	}

	printArray(&arr1)//利用指针修改数组
	fmt.Println(arr1)





}
