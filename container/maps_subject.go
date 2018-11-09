package main

import "fmt"

/*
	寻找最长不含有重复字符的子串
	例如
	abcabcbb ->  abc
    bbbbbb -> b
	pwwkew -> wke
	对于每一个字母x
	lastOccurred[x]不存在，或者 < start ->无需操作
	lastOccurred[x]存在， >=start ->更新start
	更新lastOccurred[x] ，更新maxLength
 */

func myfunc(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,ch :=  range []byte(s){
		lastI,ok := lastOccurred[ch]

		fmt.Println(lastI,ok,ch)

		if ok && lastI >= start{
			start  = lastI + 1
		}
		if i-start+1 > maxLength{
			maxLength = i -start + 1
		}
		lastOccurred[ch] = i
	}

	fmt.Println(lastOccurred)
	return maxLength

}


func myfuncall(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch :=  range []rune(s){
		lastI,ok := lastOccurred[ch]

		fmt.Println(lastI,ok,ch)

		if ok && lastI >= start{
			start  = lastI + 1
		}
		if i-start+1 > maxLength{
			maxLength = i -start + 1
		}
		lastOccurred[ch] = i
	}

	fmt.Println(lastOccurred)
	return maxLength

}


func main() {
	fmt.Println(myfunc("abcabcbb"))
	fmt.Println(myfuncall("一二三四一二三一二一"))
}
