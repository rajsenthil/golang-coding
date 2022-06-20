package algorithms

import "log"

// var prev *Node
// var next *Node

/**
Save the start. Start is nothing but HEAD for the first time.
Iterate from start to Kth item
if there is more than the Kth term (i.e Kth term is not nil)
break the link to the next item
call recursively with start+1.
if at any point the of the iteration, the loop breaks (with nil), then just return the start.
This start will be tail for the last iteration.
Now start working on the last iteration to reverse and when done add a link to recursive returned element
**/

func reverse_k_nodes(curr_head *Node, k int) *Node {
	if curr_head == nil {
		return nil
	}
	curr := curr_head
	for idx := 0; idx < k && curr != nil; idx++ {
		log.Printf("idx: %d, curr: %d", idx, curr.val)
		curr = curr.next
	}
	next_head := reverse_k_nodes(curr, k)
	if curr_head != nil {
		log.Printf("Current Head: %d", curr_head.val)
	}
	if next_head != nil {
		log.Printf("Next head: %d", next_head.val)
	}
	return reverse_k_nodes_helper(curr_head, next_head, k)
}

func reverse_k_nodes_helper(curr_head *Node, next_head *Node, k int) *Node {
	curr := curr_head
	for idx := 0; idx < k; idx++ {
		if curr == nil {
			return curr_head
		}
		curr = curr.next
	}

	curr = curr_head
	for idx := 0; idx < k; idx++ {
		if curr == nil {
			break
		}
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
		if curr != nil {
			log.Printf("reverse_k_nodes_helper loop end::: curr node: %d", curr.val)
		}
	}
	curr_head.next = next_head
	log.Printf("reverse_k_nodes_helper printing links:::")
	prev.print()
	return prev
}
