package wordmap

type Wordmap struct {
	m map[string]struct{}
}

func New() *Wordmap {
	return &Wordmap{m: make(map[string]struct{})}
}

func (m *Wordmap) Insert(word string) {
	m.m[word] = struct{}{}
}

func (m *Wordmap) Find(word string) bool {
	_, found := m.m[word]
	return found
}

func (m *Wordmap) Search(word string, distance int) map[int][]string {
	return make(map[int][]string)
}

func (m *Wordmap) List() []string {
	list := make([]string, 0, len(m.m))
	for word := range m.m {
		list = append(list, word)
	}
	return list
}
