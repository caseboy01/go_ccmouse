package strings

import (
	"fmt"
	"strings"
)

/*
	fields,split,join
	contains,index
    tolower,toupper
	trim,trimright,trimleft
*/
func main() {
	s := "abcdefgh"
	s1 := strings.Fields("1")
	s2 := strings.Split(s, "")
	s3 := strings.Join(s2, "i")

	s4 := strings.Contains(s, "tt") //是否包含

	s5 := strings.Index(s, "d")

	fmt.Println(s, s1, s2, s3, s4, s5)

	s6 := strings.ToLower("ABCEDE")
	fmt.Println(s, s6)

	s7 := strings.ToUpper(s)
	fmt.Println(s, s7)

}
