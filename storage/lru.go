package storage

import (
	"fmt"
	"fredis/types"
)

/*

SOLUTION

	we have a hashmap where we are adding and deleting values (the global map)
	we have a DLL as well
		- front indicates MRU
		- last indicates LRU

	as we start sending requests to the server, for each request we:
		- check whether it exists in the hashmap
			- if not
				- check whether we have are at CAPACITY
					- if not
						- then we create it and then add the element to the front of the DLL
						- then we store a pointer to that node in our Item that the key points to
							- the point of this is that now we know where the item currently is
					- if yes
						- then we remove the tail node and the entry from the hashmap
						- then we create it and then add the element to the front of the DLL
						- then we store a pointer to that node in our Item that the key points to
							- the point of this is that now we know where the item currently is
			- if yes
				- we check Item.Node to see where it is in the DLL
				- we remove the node and then move it to the front of the DLL

*/

const DEFAULT_MAX = 5

func removeDLL(n *types.Node) *types.Node {

	dll := GetDLL()

	prev := n.Prev
	next := n.Next

	if next != nil {
		n.Next.Prev = prev
	}

	if prev != nil {
		n.Prev.Next = next
	}

	if dll.Head == n {
		dll.Head = next
	}
	if dll.Tail == n {
		dll.Tail = prev
	}

	if prev == nil && next == nil {
		dll.Head = nil
		dll.Tail = nil
	}

	n.Next = nil
	n.Prev = nil

	return n

}

// function to remove from hashmap + DLL
func Remove(k string, hard bool) {

	cache := GetCache()

	// remove from DLL
	node := cache[k].Place

	removeDLL(node)

	if hard {
		delete(cache, k)
		dll.Length -= 1
	}

	PrintDLL()

}

func Promote(k string, v interface{}) {

	cache := GetCache()
	dll := GetDLL()

	// - check whether it exists in the hashmap
	val, exists := cache[k]

	if exists {

		node := val.Place

		// remove from its current spot if it's not already at head
		if dll.Head != node {
			newNode := removeDLL(node)

			// make it the new head
			newNode.Next = dll.Head
			newNode.Prev = nil

			if dll.Head != nil {
				dll.Head.Prev = newNode
			}

			dll.Head = newNode

			// update pointer in the cache map
			val.Place = dll.Head
			cache[k] = val
		}

	} else {

		if dll.Length == DEFAULT_MAX {
			Remove(dll.Tail.Key, true)
		}

		// create new node
		node := &types.Node{
			Key:  k,
			Prev: nil,
			Next: dll.Head,
		}

		// new head
		if dll.Head != nil {
			dll.Head.Prev = node
		}

		dll.Head = node
		if dll.Length == 0 {
			dll.Tail = dll.Head
		}

		// create new item
		val = types.Item{
			Value: v,
			Place: dll.Head,
		}

		// update pointer
		cache[k] = val

		// update length
		dll.Length += 1

	}

	PrintDLL()
}

func PrintDLL() {
	dll := GetDLL()

	curr := dll.Head
	if curr == nil {
		fmt.Println("Doubly Linked List is empty.")
		return
	}

	fmt.Print("Doubly Linked List: ")
	for curr != nil {
		fmt.Printf("[%s]", curr.Key)
		if curr.Next != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.Next
	}
	fmt.Println()
	fmt.Println("Length", dll.Length)

}
