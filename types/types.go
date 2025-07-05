package types

type Item struct {
	Value interface{} // defined as interface{} so it can be any type
	Place int         // determine order
	Time  int         // ttl, set
}

type RequestBody struct {
	Value string
}
