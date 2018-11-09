package base

import "fmt"

func main() {

	cache := make(map[string]string)
	cache["name"] = "ccmouse"
	cache["age"] = "18"

	fmt.Println(cache)
}
