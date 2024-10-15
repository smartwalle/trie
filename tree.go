package trie

type Tree[T any] struct {
	root  *Node[T]
	empty T
}

func New[T any]() *Tree[T] {
	var r = &Tree[T]{}
	r.root = newNode[T]()
	return r
}

func (t *Tree[T]) Add(key string, value T) {
	var node = t.root
	var runes = []rune(key)
	for _, r := range runes {
		if _, ok := node.children[r]; !ok {
			node.children[r] = newNode[T]()
		}
		node = node.children[r]
	}
	node.end = true
	node.value = value
}

func (t *Tree[T]) Match(key string) (bool, T) {
	var node = t.root
	var runes = []rune(key)
	for _, r := range runes {
		if _, ok := node.children[r]; !ok {
			return false, t.empty
		}
		node = node.children[r]
	}
	return node.end, node.value
}
