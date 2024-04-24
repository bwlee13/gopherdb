package lru

type Node struct {
	Key       string
	Value     interface{}
	TTL       int64
	CreatedAt int64
	Prev      *Node
	Next      *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int32
}
