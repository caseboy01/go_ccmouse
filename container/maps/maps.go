package maps

import "fmt"

/*
	map使用哈希表，必须可以比较相等
	除了slice，map，function的内建类型都可以作为key
	struct类型不包含上述字段，也可以作为key
*/

func main() {
	m := map[string]string{
		"name": "action",
		"age":  "18",
		"site": "xingdong365",
		"sex":  "boy",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	//遍历,map里的内容是无序的，遍历的内容顺序可能不同
	fmt.Println("遍历map:")
	for v, k := range m {
		fmt.Println(k, "-", v)
	}
	//获取其中一个值,如果不存在，打印的是空值
	name, ok := m["name"]
	fmt.Println(name, ok)

	//删除
	delete(m, name)
	fmt.Println(m)

}
