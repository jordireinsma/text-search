package trie

func min(xs ...int) int {
	res := xs[0]
	for _, x := range xs {
		if x < res {
			res = x
		}
	}
	return res
}

type Trie struct {
	next map[rune]*Trie
	word bool
}

func New() *Trie {
	return &Trie{next: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	ptr := t
	for _, c := range word {
		if ptr.next[c] == nil {
			ptr.next[c] = New()
		}
		ptr = ptr.next[c]
	}
	ptr.word = true
}

func (t *Trie) Find(word string) bool {
	ptr := t
	for _, c := range word {
		ptr = ptr.next[c]
		if ptr == nil {
			return false
		}
	}
	return ptr.word
}

func (t *Trie) Search(word string, distance int) map[int][]string {
	rowlen := len([]rune(word)) + 1
	rows := [2][]int{
		make([]int, rowlen),
		make([]int, rowlen),
	}
	for i := 0; i < len(rows[0]); i++ {
		rows[0][i] = i
	}
	res := make(map[int][]string)

	for c, ptr := range t.next {
		ptr.search(string(c), word, distance, c, rows, res)
	}
	return res
}

func (t *Trie) search(w, word string, distance int, c rune, rows [2][]int, res map[int][]string) {
	rows[1][0] = rows[0][0] + 1

	i := 0
	for _, cc := range word {
		insert := rows[1][i] + 1
		delete := rows[0][i+1] + 1
		replace := rows[0][i]
		if c != cc {
			replace++
		}
		rows[1][i+1] = min(insert, delete, replace)
		i++
	}
	cost := rows[1][i]
	if cost <= distance && t.word {
		res[cost] = append(res[cost], w)
	}
	if min(rows[1]...) <= distance {
		copy(rows[0], rows[1])
		for c, ptr := range t.next {
			ptr.search(w+string(c), word, distance, c, rows, res)
		}
	}
}

func (t *Trie) List() []string {
	words := []string{}
	t.list(&words, "")
	return words
}

func (t *Trie) list(words *[]string, word string) {
	if t.word {
		*words = append(*words, word)
	}
	for c, ptr := range t.next {
		ptr.list(words, word+string(c))
	}
}
