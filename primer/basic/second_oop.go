package basic

import "fmt"

// TreeNode /*   type 名字 struct {}  */
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//type 名字 struct {}
// 以.的方式调用    值接收者（go 特有的）
func (node TreeNode) print1() {
	fmt.Println(node.Value)
}

// 传入treeNode参数的形式调用
func print2(node TreeNode) {
	fmt.Println(node.Value)
}

// 指针接收者
func (node *TreeNode) setValue(value int) {
	node.Value = value
}

func setValue(node *TreeNode, value int) {
	node.Value = value
}

func RunOpp() {
	var node = TreeNode{Value: 999}
	node.print1()
	print2(node)
	node.setValue(89)
	setValue(&node, 87)
	node.print1()
}
