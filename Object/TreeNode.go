package Queue

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node *treeNode) print() {
	fmt.Print(node.value)
}

//golang没有构造方法，但可以用普通的工厂方法代替
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	node := createNode(6)
	fmt.Println(&node)
	fmt.Println(nodes)
	fmt.Println(root)
	fmt.Println()
	fmt.Println("------------")
	root.print()
}
