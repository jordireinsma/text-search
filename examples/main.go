package main

import (
	"fmt"
	"trie"
)

func main() {
	root := trie.New()
	fmt.Println("bob", trie.Find(root, "bob"))

	trie.Insert(root, "bob", "alice", "bobina", "yes")
	fmt.Println("bob", trie.Find(root, "bob"))
	fmt.Println("bobi", trie.Find(root, "bobi"))
	fmt.Println("bobina", trie.Find(root, "bobina"))
	fmt.Println("bobinas", trie.Find(root, "bobinas"))

	fmt.Println("alice", trie.Find(root, "alice"))
	trie.Insert(root, "bobi")

	fmt.Println("bobi", trie.Find(root, "bobi"))

	fmt.Println(root)

	trie.Insert(root, "")
	fmt.Println(root)
}