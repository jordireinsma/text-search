package search

// import (
// 	"math/rand"
// 	"search/trie"
// 	"search/wordmap"
// 	"testing"
// 	"unsafe"
// )
// const (
// 	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//     letterIdxBits = 6                    // 6 bits to represent a letter index
//     letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
//     letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
// )

// const size = 100000

// func randString(n int) string {
//     b := make([]byte, n)
//     // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
//     for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
//         if remain == 0 {
//             cache, remain = rand.Int63(), letterIdxMax
//         }
//         if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
//             b[i] = letterBytes[idx]
//             i--
//         }
//         cache >>= letterIdxBits
//         remain--
//     }

//     return *(*string)(unsafe.Pointer(&b))
// }

// func BenchmarkInsertTrie(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         obj := trie.New()
//         obj.Insert(words...)
//     }
// }

// func BenchmarkInsertWordMap(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         obj := wordmap.New()()
//         obj.Insert(words...)
//     }
// }

// func BenchmarkFindTrie(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := trie.New()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         for i := 0; i < size; i++ {
//             obj.Find(words[i])
//         }
//     }
// }

// func BenchmarkFindWordMap(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := wordmap.New()()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         for i := 0; i < size; i++ {
//             obj.Find(words[i])
//         }
//     }
// }

// func BenchmarkFindSimilarTrie(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := trie.New()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         for i := 0; i < size/10000; i++ {
//             obj.FindSimilar(words[i], 2)
//         }
//     }
// }

// func BenchmarkFindSimilarWordMap(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := wordmap.New()()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         for i := 0; i < size/10000; i++ {
//             obj.FindSimilar(words[i], 2)
//         }
//     }
// }

// func BenchmarkListTrie(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := trie.New()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         obj.List()
//     }
// }

// func BenchmarkListWordMap(b *testing.B) {
// 	words := []string{}
// 	for i := 0; i < size; i++ {
// 		words = append(words, randString(i % 12))
// 	}
// 	obj := wordmap.New()()
//     obj.Insert(words...)

// 	b.ResetTimer()
//     for j := 0; j < b.N; j++ {
//         obj.List()
//     }
// }


