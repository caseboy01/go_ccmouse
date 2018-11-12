package tree

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

func (root *TreeNode) traverse() {
	if root == nil {
		return
	}
	root.Left.traverse()
	fmt.Print(root.Value)
	root.Right.traverse()
}

func main() {

}
