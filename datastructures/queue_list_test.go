package datastructures

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	node5 := QueueNode{val: "5"}
	node4 := QueueNode{val: "4"}
	node3 := QueueNode{val: "3"}
	node2 := QueueNode{val: "2"}
	node1 := QueueNode{val: "1"}
	head := QueueNode{val: "0"}
	q := Queue{}
	q.queue(&head)
	q.queue(&node1)
	q.queue(&node2)
	q.queue(&node3)
	q.queue(&node4)
	q.queue(&node5)
	log.Printf("Queue: %v", q.String())
}

func TestDequeue(t *testing.T) {
	node5 := QueueNode{val: "5"}
	node4 := QueueNode{val: "4"}
	node3 := QueueNode{val: "3"}
	node2 := QueueNode{val: "2"}
	node1 := QueueNode{val: "1"}
	head := QueueNode{val: "0"}
	q := Queue{}
	q.queue(&head)
	q.queue(&node1)
	q.queue(&node2)
	q.queue(&node3)
	q.queue(&node4)
	q.queue(&node5)

	for i := 0; i <= 10; i++ {
		curr, err := q.dequeue()
		if err == nil {
			log.Printf("Dequeued Node: %v", curr.String())
		} else {
			log.Printf("Error stopped: %v", err)
			break
		}
	}
}
