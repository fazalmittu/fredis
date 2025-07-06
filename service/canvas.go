package service

/*

wanna make an default LRU both
- # of KV pairs
- any queue instantiated

default length of 100

general cache

have a global array of keys
- first item is MRU, last is LRU

set
	creates or updates key in map. O(1)
	moves to top of array O(n)

get
	retrieve item from map. O(1)
	moves to top of array. O(n)

delete
	deletes item from map and array O(n)


linked list
	add O(1)
	retrieve O(n)
	update O(n)

	add to arbitrary point O(n)

hashmap
	add O(1)
	retrieve O(1)
	update O(1)

	traverse keys O(n)

array
	add to end O(1)
	add to arbitrary point O(n)
	retrieve O(1)
	updating O(1)

min heap
	add O(log(n))
	pop top O(log(n))
	heapify O(n*log(n))


-------------------------------

set
	creates or updates key in map		O(1)
	will add to min heap 				O(log(n))

get
	retrieves item from map				O(1)
	retrives from min heap				O(log(n))

delete
	deletes from map 					O(1)
	deletes item						O(log(n))


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
