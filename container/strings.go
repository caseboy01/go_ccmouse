package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(len(s)) // 19
	fmt.Println(utf8.RuneCountInString(s))// 9
	//1 遍历
	bytes := []byte(s)
	fmt.Println(bytes)
	for len(bytes) >0 {
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}
	fmt.Println(bytes)
	//2 遍历
	for i,ch := range [] rune(s){
		fmt.Printf("(%d,%c)",i,ch)
	}

}
