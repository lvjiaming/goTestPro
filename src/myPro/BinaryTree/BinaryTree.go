/**
 二叉树的接口以及结构
 */
package BinaryTree

import "fmt"

func init()  { // 包在加载时会调用此函数
	fmt.Printf("初始化")
}

type Node struct {
	Data interface{}  // 存放的数据
	Left *Node // 左边节点的地址
	Right *Node // 右边节点的地址
}

type Initer interface {
	SetData(data interface{}) // 设置数据
}

type Operater interface {
	PrintBT()  // 打印二叉树
	Depth() int  // 二叉树的深度
	LeafCount() int // 获取叶子节点
}

type Order interface { // 遍历
	PreOrder() // 前序遍历
	InOrder() // 中序遍历
	PostOrder() // 后续遍历
}

func (n *Node) SetData (data interface{})  {
	n.Data = data
}

func (n *Node) PrintBT ()  {
	fmt.Printf("%v", n.Data)
	fmt.Printf("(")
	if n.Left != nil {
		n.Left.PrintBT()
	}
	if n.Right != nil {
		fmt.Printf(",")
		n.Right.PrintBT()
	}
	fmt.Printf(")")
}

func (n *Node) Depth () int {
	return Depth(n)
}

func (n *Node) LeafCount () int {
	return LeafCount(n)
}

func (n *Node) PreOrder ()  {

}

func (n *Node) InOrder ()  {

}

func (n *Node) PostOrder ()  {

}

func NewNode() *Node {
	return &Node{
		Data:  nil,
		Left:  nil,
		Right: nil,
	}
}

func Depth(node *Node) int {
	if node == nil {
		return 0
	}
	var leftDepth, rightDepth int
	leftDepth = Depth(node.Left)
	rightDepth = Depth(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func LeafCount(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	} else {
		return LeafCount(node.Left) + LeafCount(node.Right)
	}
}





