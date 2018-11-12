package tree

import "fmt"

func (root *Node) Traverse() {
	if root == nil {
		return
	}
	root.Left.Traverse()
	fmt.Print(root.Value, "\n")
	root.Right.Traverse()
}
