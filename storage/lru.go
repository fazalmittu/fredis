package storage

import (
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

const DEFAULT_MAX = 100

func removeDLL(n *types.Node) *types.Node {

	dll := GetDLL()

	prev := n.Prev
	next := n.Next

	if next != nil {
		n.Next.Prev = prev
		dll.Tail = prev
	}

	if prev != nil {
		n.Prev.Next = next
		dll.Head = next
	}

	dll.Length -= 1

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
	}

}

func Promote(k string) {

	cache := GetCache()

	// - check whether it exists in the hashmap
	val, exists := cache[k]

	if exists {

		// remove from its current spot
		node := val.Place
		newNode := removeDLL(node)

		dll := GetDLL()

		// make it the new head
		newNode.Next = dll.Head
		dll.Head = newNode

		// update pointer in the cache map
		val.Place = dll.Head
		cache[k] = val

	} else {

		dll := GetDLL()

		if dll.Length == DEFAULT_MAX {
			removeDLL(dll.Tail)
		}

		// create new node
		node := types.Node{
			Key:  k,
			Prev: nil,
			Next: dll.Head,
		}

		// new head
		dll.Head = &node
		if dll.Length == 0 {
			dll.Tail = dll.Head
		}

		// update pointer
		val.Place = dll.Head
		cache[k] = val

		// update length
		dll.Length += 1

	}

}
