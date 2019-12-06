package main

import "fmt"

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// NodeBetweenValues method of LinkedList class returns the node that has a property lying
// between the firstProperty and secondProperty values.
func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead method of LinkedList
/**
The AddToHead method of the doubly LinkedList class sets the previousNode property of
the current headNode of the linked list to the node that's added with property.
The node with property will be set as the headNode of the LinkedList method in the following code:
*/
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		// fmt.Println(node.property)
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

// NodeWithValue method of LinkedList
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method of LinkedList
/**
The AddAfter method adds a node after a specific node to a double linked list.
The AddAfter method of the double LinkedList class searches the node whose value is equal to nodeProperty.
The found node is set as the previousNode of the node that was added with property.
*/
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

/**
The AddToEnd method adds the node to the end of the double linked list.
The AddToEnd of the LinkedList class creates a node whose property is
set as the integer parameter property.
*/
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	fmt.Println(linkedList.headNode.property)

	var node *Node
	node = linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
