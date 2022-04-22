package datastructures

import (
	"fmt"
	"log"
)

type BinaryTreeArrayNode struct {
	val        int
	leftChild  *BinaryTreeArrayNode
	rightChild *BinaryTreeArrayNode
}

type BinaryTreeArray struct {
	root *BinaryTreeArrayNode
}

func (n *BinaryTreeArrayNode) String() string {
	s := ""
	if n != nil {
		val := n.val
		left := n.leftChild
		right := n.rightChild
		s += left.String() + fmt.Sprintf("Node: %d\t", val) + right.String()
		return s
	}
	return s
}

func (t *BinaryTreeArray) String() string {
	log.Printf("Tree String generation...")
	return t.root.String()
}

func (t *BinaryTreeArray) addToBinaryTree(val int) *BinaryTreeArray {
	log.Printf("First line addToBinaryTree")
	currval := t.root
	log.Printf("Entering addToBinaryTree: %v", currval)
	if currval == nil {
		root := BinaryTreeArrayNode{val: val}
		t.root = &root
		log.Printf("Setting the root: %v", root)
		log.Printf("Returning the tree: %v", t)
		return t
	}

	parent := currval
	log.Printf("Entering after root: %v", currval)
	for currval != nil {
		parent = currval
		log.Printf("Prev Node: %v, Current Node: %v", parent, currval.val)
		if val <= currval.val {
			currval = currval.leftChild
		} else {
			currval = currval.rightChild
		}
	}
	newNode := &BinaryTreeArrayNode{val: val}
	if val <= parent.val {
		parent.leftChild = newNode
	} else {
		parent.rightChild = newNode
	}
	return t
}

func (n *BinaryTreeArrayNode) preorder() string {
	s := ""
	if n == nil {
		log.Print("The preorder is called with nil")
		return ""
	}
	log.Printf("Pre-Order traversal starting: %d\n", n.val)
	currNode := n
	leftChild := currNode.leftChild
	s += leftChild.preorder()
	s += fmt.Sprintf(" %d ", currNode.val)
	rightChild := currNode.rightChild
	s += rightChild.preorder()
	log.Printf("Pre-Order traversal: %s", s)
	return s
}
