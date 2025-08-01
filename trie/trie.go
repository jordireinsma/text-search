package trie

type Trie struct {
	next  map[rune]*Trie
	final bool
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
	ptr.final = true
}

func (t *Trie) Find(word string) bool {
	ptr := t
	for _, c := range word {
		ptr = ptr.next[c]
		if ptr == nil {
			return false
		}
	}
	return ptr.final
}

type edits struct {
	target   []rune
	distance int
}

func (t *Trie) Fuzzy(word string, distance int) map[int][]string {
	rowlen := len([]rune(word)) + 1
	rows := [3][]int{
		make([]int, rowlen), // previous row
		make([]int, rowlen), // current row
		make([]int, rowlen), // next row
	}
	for i := 0; i < rowlen; i++ {
		rows[0][i] = rowlen
		rows[1][i] = i
	}
	e := edits{
		target:   []rune(word),
		distance: distance,
	}
	res := make(map[int][]string)
	for c, next := range t.next {
		e.run(next, []rune{c}, rows, res)
	}
	return res
}

func (e *edits) run(ptr *Trie, word []rune, rows [3][]int, res map[int][]string) {
	rows[2][0] = rows[1][0] + 1
	i := 0
	for ; i <= len(e.target); i++ {
		insert := rows[2][i] + 1
		delete := rows[1][i+1] + 1
		replace := rows[1][i]
		transpose := insert
		if word[len(word)-1] != e.target[i] {
			replace++
			if i > 1 && len(word) > 1 && word[len(word)-2] == e.target[i] && word[len(word)-1] == e.target[i-1] {
				transpose = rows[0][i-1] + 1
			}
		}
		rows[2][i+1] = min(insert, delete, replace, transpose)
	}
	cost := rows[2][i]
	if cost <= e.distance && ptr.final {
		res[cost] = append(res[cost], string(word))
	}
	if min(rows[2]...) <= e.distance {
		copy(rows[0], rows[1])
		copy(rows[1], rows[2])
		for c, next := range ptr.next {
			e.run(next, append(word, c), rows, res)
		}
	}
}

func (t *Trie) List() []string {
	words := []string{}
	t.list(&words, "")
	return words
}

func (t *Trie) list(words *[]string, word string) {
	if t.final {
		*words = append(*words, word)
	}
	for c, next := range t.next {
		next.list(words, word+string(c))
	}
}

func min(xs ...int) int {
	res := xs[0]
	for _, x := range xs {
		if x < res {
			res = x
		}
	}
	return res
}
