package types

type Item struct {
	Key   string
	Value interface{} // defined as interface{} so it can be any type
	Place int         // determine order
	Time  int         // ttl, set
}
