package main

import (
	"fmt"

	"github.com/jordi-reinsma/text-search/trie"
)

func main() {
	words := []string{
		"bob", "alice", "bobina", "bo",
	}

	root := trie.New()
	fmt.Println("bob", root.Find("bob"))

	for _, word := range words {
		root.Insert(word)
	}

	for _, word := range words {
		fmt.Println(word, root.Find(word))
	}

	fmt.Println("all", root.Find(""))
	fmt.Println(root.List())

	fmt.Println(root.Search("bob", 1))
	fmt.Println(root.Search("boba", 1))
	fmt.Println(root.Search("bobi", 2))

	root.Insert("")
	fmt.Println("all", root.Find(""))
	fmt.Println(root.List())
}
