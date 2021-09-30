package main

import (
	"fmt"
)

// define the node for doubly linked list
type node struct {
	prev, next *node
	data       interface{}
}

// init the head node.
var list node

// // init the value for head node.
func init() {
	list.data = "HeadNode"
	list.next, list.prev = nil, nil
}

// insert a node which value form arg to the fist for list.
func insert(data ...interface{}) {
	for _, x := range data {
		tempNode := node{prev: &list, next: list.next, data: x}
		if list.next != nil {
			list.next.prev = &tempNode
		}
		list.next = &tempNode
	}
}

// delete a node from list by given node.
func delete(n *node) {
	if n == &list {
		println("cannot remove head node!")
		return
	}
	n.prev.next = n.next
}

// deleteFirst delete the fist node of the list.
func deleteFirst() {
	delete(list.next)
}

// deleteLast delete the last node of the list.
func deleteLast() {
	var lastNode node
	lastNode.next = list.next

	for lastNode.next != nil {
		lastNode = *lastNode.next
	}
	delete(&lastNode)
}

// deleteData delete the node by given node's data.
func deleteData(data interface{}) {
	delete(listSearch(data))
}

// printList print the linklist.
// the headNode is the head node for the list.
func printList(headNode node) {
	if headNode.next == nil {
		println("list is empty!")
		return
	}
	fmt.Println("________________________")
	fmt.Println("0", "\t", headNode.data /* , "\t", &headNode.prev, "\t", &headNode.next */)
	var lenth int
	for headNode.next != nil {
		lenth++
		fmt.Println(lenth, "\t", headNode.next.data /*, "\t"  , &headNode.next.prev, "\t", &headNode.next.next */)
		headNode = *headNode.next
	}
}

// listSearch return a node by given x from list.
// the x is data of node.
func listSearch(x interface{}) *node {
	var pointer node
	pointer = list
	for pointer.data != x {
		if pointer.next == nil {
			println("no node found!")
			return nil
		}
		pointer = *pointer.next
	}
	return &pointer
}

func main() {
	insert("第一", 123, "ssss", []int{3, 4, 654, 256, 456, 4564}, "sun", "two")
	printList(list)

	deleteData(123)
	printList(list)

	/*
		________________________
		0        HeadNode
		1        two
		2        sun
		3        [3 4 654 256 456 4564]
		4        ssss
		5        123
		6        第一
		________________________
		0        HeadNode
		1        two
		2        sun
		3        [3 4 654 256 456 4564]
		4        ssss
		5        第一

	*/
}
