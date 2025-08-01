package search

type Dictionary interface {
	Insert(word string)
	Find(word string) bool
	Fuzzy(word string, distance int) map[int][]string
	List() []string
}
