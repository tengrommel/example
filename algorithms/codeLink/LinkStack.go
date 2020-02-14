package main

type Node struct {
	data  interface{}
	pNext *Node
}

func NewStack() *Node {
	return &Node{} // 返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newNode := &Node{data: data}
	newNode.pNext = n.pNext
	n.pNext = newNode
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.pNext.data // 要弹出的数据
	n.pNext = n.pNext.pNext
	return value
}

func (n *Node) Length() int {
	pNext := n
	length := 0
	for pNext.pNext != nil {
		pNext = pNext.pNext // 节点循环
		length++
	}
	return length
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}
