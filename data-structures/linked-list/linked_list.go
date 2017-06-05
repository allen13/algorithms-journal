package linkedlist

//Node - LinkedList primative type
type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

//LinkedList - basis for linked list
type LinkedList struct {
	head *Node
	tail *Node
}

//IntAt - returns int value at given index
func (ll *LinkedList) IntAt(index int) int {
	currentNode := ll.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value.(int)
}

//PushInt - add int value to list
func (ll *LinkedList) PushInt(value int) {
	node := &Node{value: value}

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}
}
