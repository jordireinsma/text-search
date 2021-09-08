package trie

import (
	"strings"
)

type Node struct {
	ptr map[rune]*Node
	word bool
}

func New() *Node {
	return &Node{ptr:make(map[rune]*Node)}
}

func Insert(t *Node, words ...string) {
	for _, word := range words {
		it := t
		for _, c := range word {
			if it.ptr[c] == nil {
				it.ptr[c] = New()
			}
			it = it.ptr[c]
		}
		it.word = true
	}
}

func Find(t *Node, word string) bool {
	it := t
	for _, c := range word {
		if it.ptr[c] == nil {
			return false
		}
		it = it.ptr[c]
	}
	return it.word
}

// func (t *Node) FindSimilar(word string, distance int) string {
// 	res := ""
// 	it := t
// 	for _, c := range word {
// 		if it.ptr[c] == nil {
// 			return 
// 		}
// 		it = it.ptr[c]
// 	}
// 	return it.word
// }

func List(t *Node) []string {
	return list(t, "")
}

func list(t *Node, word string) []string {
	words := []string{}
	if t.word {
		words = append(words, word)
	}
	for c, it := range t.ptr {
		tmp := list(it, word + string(c))
		words = append(words, tmp...)
	}
	return words
}

func (t *Node) String() string {
	return strings.Join(list(t, ""), ", ")
}
