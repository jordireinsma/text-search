package wordmap

type Wordmap struct {
	words    map[string]struct{}
	alphabet map[rune]struct{}
}

func New() *Wordmap {
	return &Wordmap{
		words:    make(map[string]struct{}),
		alphabet: make(map[rune]struct{}),
	}
}

func (m *Wordmap) Insert(word string) {
	for _, c := range word {
		m.alphabet[c] = struct{}{}
	}
	m.words[word] = struct{}{}
}

func (m *Wordmap) Find(word string) bool {
	_, found := m.words[word]
	return found
}

func (m *Wordmap) Fuzzy(word string, distance int) map[int][]string {
	res := make(map[int][]string)
	words := make(map[string]bool)
	if _, found := m.words[word]; found && distance >= 0 {
		res[0] = []string{word}
		words[word] = true
	} else {
		words[word] = false
	}
	alphabet := make([]string, 0, len(m.alphabet))
	for c := range m.alphabet {
		alphabet = append(alphabet, string(c))
	}
	for d := 1; d <= distance; d++ {
		words = edits(alphabet, words)
		for w := range words {
			if visited := words[w]; visited {
				continue
			}
			if _, found := m.words[w]; found {
				res[d] = append(res[d], w)
			}
			words[w] = true // visited
		}
	}
	return res
}

func edits(alphabet []string, words map[string]bool) map[string]bool {
	variants := make([]string, 2*len(words)*len(alphabet))
	for word := range words {
		for i := 0; i < len(word); i++ {
			for _, c := range alphabet {
				tmp := word[:i] + c
				// insert
				variants = append(variants, tmp+word[i:])
				// replace
				variants = append(variants, tmp+word[i+1:])
			}
			// delete
			variants = append(variants, word[:i]+word[i+1:])
			// transpose
			if i < len(word)-1 {
				variants = append(variants, word[:i]+word[i+1:i+2]+word[i:i+1]+word[i+2:])
			}
		}
		// final insert
		for _, c := range alphabet {
			variants = append(variants, word+c)
		}
	}
	res := make(map[string]bool, 2*len(words))
	for _, variant := range variants {
		res[variant] = words[variant]
	}
	return res
}

func (m *Wordmap) List() []string {
	list := make([]string, 0, len(m.words))
	for word := range m.words {
		list = append(list, word)
	}
	return list
}
