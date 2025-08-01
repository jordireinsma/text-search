package search

import (
	"math/rand"
	"testing"
	"unsafe"

	"github.com/jordi-reinsma/text-search/trie"
	"github.com/jordi-reinsma/text-search/wordmap"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

const inserts = 100000
const wordlen = 8

func randString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func insert(d Dictionary, b *testing.B) {
	words := make([]string, inserts)
	for i := 0; i < inserts; i++ {
		words[i] = randString(i % wordlen)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		obj := d
		for _, word := range words {
			obj.Insert(word)
		}
	}
}

func BenchmarkInsertTrie(b *testing.B) {
	insert(trie.New(), b)
}

func BenchmarkInsertWordMap(b *testing.B) {
	insert(wordmap.New(), b)
}

func find(d Dictionary, b *testing.B) {
	words := make([]string, inserts)
	for i := 0; i < inserts; i++ {
		words[i] = randString(i % wordlen)
	}
	obj := d
	for _, word := range words {
		obj.Insert(word)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		obj.Find(words[rand.Int63n(inserts)])
	}
}

func BenchmarkFindTrie(b *testing.B) {
	find(trie.New(), b)
}

func BenchmarkFindWordMap(b *testing.B) {
	find(wordmap.New(), b)
}

func search(d Dictionary, b *testing.B) {
	words := make([]string, inserts)
	for i := 0; i < inserts; i++ {
		words[i] = randString(i % wordlen)
	}
	obj := d
	for _, word := range words {
		obj.Insert(word)
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		obj.Fuzzy(words[rand.Int63n(inserts)], 2)
	}
}

func BenchmarkFuzzyTrie(b *testing.B) {
	search(trie.New(), b)
}

func BenchmarkFuzzyWordMap(b *testing.B) {
	search(wordmap.New(), b)
}

func list(d Dictionary, b *testing.B) {
	words := make([]string, inserts)
	for i := 0; i < inserts; i++ {
		words[i] = randString(i % wordlen)
	}
	obj := d
	for _, word := range words {
		obj.Insert(word)
	}

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		obj.List()
	}
}

func BenchmarkListTrie(b *testing.B) {
	list(trie.New(), b)
}

func BenchmarkListWordMap(b *testing.B) {
	list(wordmap.New(), b)
}
