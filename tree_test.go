package trie_test

import (
	"github.com/smartwalle/trie"
	"testing"
)

func TestTree_Match(t *testing.T) {
	var tree = trie.New[string]()
	tree.Add("10", "v10")
	tree.Add("11", "v11")
	tree.Add("12", "v12")
	tree.Add("123", "v123")

	t.Log(tree.Match("1"))
	t.Log(tree.Match("10"))
	t.Log(tree.Match("11"))
	t.Log(tree.Match("13"))
	t.Log(tree.Match("123"))
	t.Log(tree.Match("1234"))
}
