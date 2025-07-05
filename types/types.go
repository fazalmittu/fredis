package types

type Item struct {
	Key   string
	Value interface{} // defined as interface{} so it can be any type
	Time  int         // ttl, set
}
