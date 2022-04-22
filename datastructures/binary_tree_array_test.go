package datastructures

import (
  "log"
  "testing"
)

func TestBinaryTreeArray(t *testing.T) {
  tr := BinaryTreeArray{}
  log.Printf("Setting the root with value: 1")
  tr.addToBinaryTree(1)
  log.Printf("Set to the root with value: 1")
  log.Printf("Setting the root with value: 2")
  tr.addToBinaryTree(2)
  log.Printf("Set to the root with value: 2")
  log.Printf("Setting the root with value: 3")
  tr.addToBinaryTree(3)
  log.Printf("Set to the root with value: 3")
  log.Printf("Setting the root with value: 4")
  tr.addToBinaryTree(4)
  tr.addToBinaryTree(5)
  tr.addToBinaryTree(6)
  tr.addToBinaryTree(7)

  log.Printf("Printing the tree")
  log.Printf("Tree:= %s", tr.String())
}

func TestBinaryTreeArrayPreOrder(t *testing.T) {
  tr := BinaryTreeArray{}
  tr.addToBinaryTree(4)
  log.Printf("Setting the root with value: 1")
  tr.addToBinaryTree(1)
  log.Printf("Set to the root with value: 1")
  log.Printf("Setting the root with value: 2")
  tr.addToBinaryTree(5)
  tr.addToBinaryTree(2)
  log.Printf("Set to the root with value: 2")
  log.Printf("Setting the root with value: 3")
  tr.addToBinaryTree(7)
  tr.addToBinaryTree(6)
  tr.addToBinaryTree(3)
  log.Printf("Set to the root with value: 3")
  log.Printf("Setting the root with value: 4")
  root := tr.root
  log.Printf("RESULT::: Pre-Order traversal: %s", root.preorder())
}
