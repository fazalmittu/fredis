package types

type Item struct {
	Value interface{} // defined as interface{} so it can be any type
	Place int         // determine order
	Time  int         // ttl, set
}

type CoreRequestBody struct {
	Value string
}

type CounterRequestBody struct {
	Amount int
}
