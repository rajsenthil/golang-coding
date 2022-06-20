package algorithms

import (
	"log"
	"testing"
)

func TestReverseKNodes(t *testing.T) {
	node_1 := &Node{val: 1}
	node_2 := &Node{val: 2}
	node_3 := &Node{val: 3}
	node_4 := &Node{val: 4}
	node_5 := &Node{val: 5}

	node_1.next = node_2
	node_2.next = node_3
	node_3.next = node_4
	node_4.next = node_5

	k := 2
	result := reverse_k_nodes(node_1, k)
	log.Printf("Printing the results")
	result.print()
}
