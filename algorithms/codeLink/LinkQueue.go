package main

type QueueLink struct {
	rear  *Node
	front *Node
}

func (q *QueueLink) length() int {
	next := q.front
	length := 0
	for next.pNext != nil {
		next = next.pNext
		length++
	}
	return length
}

func (q *QueueLink) Enqueue(value interface{}) {
	newNode := &Node{value, nil}
	if q.front == nil {
		q.front = newNode // 插入一个节点
		q.rear = newNode
	} else {
		q.rear.pNext = newNode
		q.rear = q.rear.pNext
	}
}

func (q *QueueLink) Dequeue() interface{} {
	if q.front == nil {
		return nil
	}
	newNode := q.front     // 记录头部位置
	if q.front == q.rear { // 只剩下一个
		q.front = nil
		q.rear = nil
	} else {
		q.front = q.front.pNext
	}
	return newNode.data
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

// 节点可以共用
type LinkQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() interface{}
}
