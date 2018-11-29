package main

import (
	"fmt"
	"math"
)

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

func lengthOfNonRepeatingSubstr(s string) int {
	//lastOccurred := make(map[rune]int)
	lastOccurred := make([]int, 0xfffff) //65535 用空间换时间
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI := lastOccurred[ch]

		//fmt.Println(lastI,ok,ch)

		if lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	//fmt.Println(lastOccurred)
	return maxLength

}

func main() {
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubstr("一二三四一二三一二一"))

	tmp := math.MaxInt32 + math.MinInt32
	fmt.Println(tmp)
}
