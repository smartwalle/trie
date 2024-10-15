package trie

type Node[T any] struct {
	end      bool
	value    T
	children map[rune]*Node[T]
}

func newNode[T any]() *Node[T] {
	var n = &Node[T]{}
	n.children = make(map[rune]*Node[T])
	return n
}
