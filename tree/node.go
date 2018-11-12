package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

//上面的写法类似于这种
func Prints(node Node) {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node")
		return
	}
	node.Value = value
}

//工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
