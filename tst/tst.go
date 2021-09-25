package tst

type TST struct {
	l, m, r *TST
	word    bool
	key     rune
}

func New() *TST {
	return &TST{}
}

func (t *TST) Insert(word string) {
	ptr := t
	for _, c := range word {
		for c != ptr.key {
			switch {
			case c < ptr.key:
				if ptr.l == nil {
					ptr.l = &TST{key: c}
				}
				ptr = ptr.l
			case c > ptr.key:
				if ptr.r == nil {
					ptr.r = &TST{key: c}
				}
				ptr = ptr.r
			default:
				ptr.key = c
				if ptr.m == nil {
					ptr.m = &TST{}
				}
				ptr = ptr.m
			}
		}
	}
	ptr.word = true
}

func (t *TST) Find(word string) bool {
	ptr := t
	for _, c := range word {
		for c != ptr.key {
			switch {
			case c < ptr.key:
				ptr = ptr.l
			case c > ptr.key:
				ptr = ptr.r
			default:
				ptr = ptr.m
			}
			if ptr == nil {
				return false
			}
		}
	}
	return ptr.word
}

func (t *TST) Search(word string, distance int) map[int][]string {
	return make(map[int][]string)
}

func (t *TST) List() []string {
	return t.list("")
}

func (t *TST) list(word string) []string {
	words := []string{}
	if t.word {
		words = append(words, word)
	}
	if t.l != nil {
		words = append(words, t.l.list(word)...)
	}
	if t.m != nil {
		words = append(words, t.m.list(word+string(t.key))...)
	}
	if t.r != nil {
		words = append(words, t.r.list(word)...)
	}
	return words
}
