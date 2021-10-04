package main

import "fmt"

// Node is the unit of the binary search tree
type Node struct {
	value               int
	left, right, parent *Node
}

// insert node to tree by val
func (head *Node) insert(val int) {
	if head == nil {
		return
	}

	node := &Node{val, nil, nil, nil}
	if head.value > val {
		if head.left == nil {
			head.left, node.parent = node, head
			return
		}

		head.left.insert(val)
	} else {
		if head.right == nil {
			head.right, node.parent = node, head
			return
		}

		head.right.insert(val)
	}
}

// find node by val in the tree, return a nil pointer if node not exist
func (head *Node) find(val int) *Node {
	var pointer *Node = head
	for pointer != nil && pointer.value != val {
		if pointer.value < val {
			pointer = pointer.right
		} else {
			pointer = pointer.left
		}
	}
	return pointer
}

// delete node in the tree by value
func (head *Node) delete(val int) error {
	var pointer *Node
	if pointer = head.find(val); pointer == nil {
		return fmt.Errorf("no data found")
	}

	// pointer's children are exist
	if pointer.left != nil && pointer.right != nil {
		pointer.value = pointer.left.value
		pointer.left.delete(pointer.value)
		return nil
	}
	// pointer's children are not exist
	if pointer.left == nil && pointer.right == nil {
		if pointer.parent.left == pointer {
			pointer.parent.left, pointer.parent = nil, nil
		} else {
			pointer.parent.right, pointer.parent = nil, nil
		}
		return nil
	}

	// pointer child exist
	if pointer.left != nil {
		pointer.left.parent, pointer.parent.left = pointer.parent, pointer.left
		pointer.empty()
	} else {
		pointer.right.parent, pointer.parent.right = pointer.parent, pointer.right
		pointer.empty()
	}
	return nil
}

// empty the current node value
func (head *Node) empty() {
	head.value = 0
	head.left, head.right, head.parent = nil, nil, nil
}

//preorder print the tree
func preorder(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ", head.value)
	preorder(head.left)
	preorder(head.right)
}

//inorder print the tree
func inorder(head *Node) {
	if head == nil {
		return
	}
	inorder(head.left)
	fmt.Printf("%v ", head.value)
	inorder(head.right)
}

func main() {
	// input:30, 88, 12, 1, 20, 17, 25
	head := Node{30, nil, nil, nil}
	for _, val := range []int{88, 12, 1, 20, 17, 25} {
		head.insert(val)
	}

	// preorder:1 12 17 20 25 30 88
	print("preorder:")
	preorder(&head)
	println()

	// inorder:30 12 1 20 17 25 88
	print("inorder:")
	inorder(&head)
	println()
	// find 12: 12
	if res := head.find(12); res != nil {
		fmt.Println("find 12:", res.value)
	} else {
		fmt.Println("no date found: ", 12)
	}
	// no date found:  100
	if res := head.find(100); res != nil {
		fmt.Println("find 100:", res.value)
	} else {
		fmt.Println("no date found: ", 100)
	}

	// preorder:1 17 20 25 12 88
	// inorder:12 1 20 17 25 88
	head.delete(30)
	print("preorder:")
	preorder(&head)
	println()

	print("inorder:")
	inorder(&head)
	println()

}
