package main

import (
	"fmt"
	"mooc_google_shizhan/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //new 返回的是指针
	root.Left.Right = tree.CreateNode(2)

	fmt.Println(root)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()      // 3
	tree.Prints(root) //3

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	//var mRoot  *treeNode
	//mRoot.setValue(5)
	//mRoot.print()

	fmt.Println()
	fmt.Println(root, root.Value, root.Left.Left, root.Left.Right)
	fmt.Println()
	root.Traverse()

}
