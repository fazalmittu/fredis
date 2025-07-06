package types

type Item struct {
	Value interface{} // defined as interface{} so it can be any type
	Time  int         // ttl, set
	Place *Node
}

type Node struct {
	Key  string
	Prev *Node
	Next *Node
}

type DLL struct {
	Head   *Node
	Tail   *Node
	Length int
}

type CoreRequestBody struct {
	Value string
}

type CounterRequestBody struct {
	Amount int
}

type Queue struct {
}
