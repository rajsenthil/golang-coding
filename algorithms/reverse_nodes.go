package algorithms

import (
	"log"
	"sync"
)

type Node struct {
	lock sync.Mutex
	val  any
	next *Node
}

var prev *Node
var next *Node
var tail *Node

func reverse_nodes(head *Node) *Node {
	curr := head
	for curr != nil {
		curr.lock.Lock()
		defer curr.lock.Unlock()
		log.Printf("Curr: %d", curr.val)
		next = curr.next
		if next == nil {
			tail = curr
		}
		curr.next = prev
		prev = curr
		curr = next
	}

	log.Printf("Afer reversed Prev: %d, Prev.next: %d", prev.val, prev.next.val)
	return prev
}

func (n *Node) print() {
	for n != nil {
		log.Printf("Node val: %v", n.val)
		if n.next != nil {
			log.Printf("Node next val: %v", n.next.val)
		} else {
			log.Printf("Node next is nil")
		}
		n = n.next
	}
}
