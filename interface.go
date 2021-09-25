package search

type Dictionary interface {
	Insert(word string)
	Find(word string) bool
	Search(word string, distance int) map[int][]string
	List() []string
}
