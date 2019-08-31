package basic

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	value       int
}

type MyTreeNode struct {
	node *TreeNode
}

//传地址，会改变内容
func (node *TreeNode) printTree() {
	fmt.Print(node.value)
	node.value = 101
}

//传拷贝,不改变原值
func (node TreeNode) setValue(val int) {
	node.value = val
}

func createNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func TreeNodeDemo() {
	//var root TreeNode
	root := TreeNode{
		value: 3,
	}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{nil, nil, 6}
	root.Right.Left = new(TreeNode)
	root.Left.Right = createNode(1)

	root.printTree()
	fmt.Println("root *:", root.value)

	root.setValue(100) //并不会改变原root的value
	fmt.Println("root:", root.value)
	nodes := []TreeNode{
		{value: 3},
		{},
		{nil, nil, 2},
		{&root, nil, 9},
	}
	fmt.Println(nodes)
}

func main() {
	//4-1
	TreeNodeDemo()
	//4-2 包和封装，结构体名、函数名、feild名大小写的作用域
	//扩展包
}
