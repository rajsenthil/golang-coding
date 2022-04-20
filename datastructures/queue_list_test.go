package datastructures

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	node5 := QueueListNode{val: "5"}
	node4 := QueueListNode{val: "4"}
	node3 := QueueListNode{val: "3"}
	node2 := QueueListNode{val: "2"}
	node1 := QueueListNode{val: "1"}
	head := QueueListNode{val: "0"}
	q := QueueList{}
	q.queueList(&head)
	q.queueList(&node1)
	q.queueList(&node2)
	q.queueList(&node3)
	q.queueList(&node4)
	q.queueList(&node5)
	log.Printf("Queue: %v", q.String())
}

func TestDequeue(t *testing.T) {
	node5 := QueueListNode{val: "5"}
	node4 := QueueListNode{val: "4"}
	node3 := QueueListNode{val: "3"}
	node2 := QueueListNode{val: "2"}
	node1 := QueueListNode{val: "1"}
	head := QueueListNode{val: "0"}
	q := QueueList{}
	q.queueList(&head)
	q.queueList(&node1)
	q.queueList(&node2)
	q.queueList(&node3)
	q.queueList(&node4)
	q.queueList(&node5)

	for i := 0; i <= 10; i++ {
		curr, err := q.dequeueList()
		if err == nil {
			log.Printf("Dequeued Node: %v", curr.String())
		} else {
			log.Printf("Error stopped: %v", err)
			break
		}
	}
}
