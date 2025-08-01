package symdel

import (
	"sort"
)

type WordData map[string]interface{}

type Entry struct {
	Word string
	Data *WordData
}

type Suggestion struct {
	Entry Entry
	Score float64 // lower is better
}

type EditDistance func(Entry, Entry) float64

type SymDel struct {
	words    map[string]*WordData
	deletes  map[string][]string
	distance EditDistance
	maxDepth int
}

type Option struct {
	DistanceFunc EditDistance
	MaxDepth     int
	InitialCap   int
	_            struct{}
}

var defaultOption = Option{
	DistanceFunc: DemerauLevenshteinDistance,
	MaxDepth:     2,
	InitialCap:   16,
}

func New(opt ...Option) *SymDel {
	s := new(SymDel)
	if len(opt) == 0 {
		opt = []Option{defaultOption}
	}
	expectedDeletes := opt[0].MaxDepth * opt[0].MaxDepth * opt[0].InitialCap
	s.words = make(map[string]*WordData, opt[0].InitialCap)
	s.deletes = make(map[string][]string, expectedDeletes)
	s.distance = opt[0].DistanceFunc
	s.maxDepth = opt[0].MaxDepth
	return s
}

func (s *SymDel) Insert(entry Entry) {
	_, found := s.words[entry.Word]
	s.words[entry.Word] = entry.Data
	if found {
		return
	}
	newDeletes := CreateDeletes(entry.Word, s.maxDepth)
	for _, newDelete := range newDeletes {
		found := false
		for _, delete := range s.deletes[newDelete] {
			if delete == entry.Word {
				found = true
				break
			}
		}
		if !found {
			s.deletes[newDelete] = append(s.deletes[newDelete], entry.Word)
		}
	}
}

func (s *SymDel) Find(word string) *WordData {
	return s.words[word]
}

func (s *SymDel) Suggest(entry Entry) []Suggestion {
	target := entry.Word
	suggestions := make(map[string]Suggestion)
	// Word matches with a dictionary word
	if data, found := s.words[target]; found {
		suggestions[target] = s.newSuggestion(
			target, target, entry.Data, data,
		)
	}
	// Word matches with precomputed dictionary words
	if deletes, found := s.deletes[target]; found {
		for _, delete := range deletes {
			suggestions[delete] = s.newSuggestion(
				target, delete, entry.Data, s.words[delete],
			)
		}
	}
	targetDeletes := CreateDeletes(target, s.maxDepth)
	// Precomputed words matches with a dictionary word
	for _, targetDelete := range targetDeletes {
		data, found := s.words[targetDelete]
		if !found {
			continue
		}
		suggestion := s.newSuggestion(targetDelete, target, data, entry.Data)
		currentSuggestion, found := suggestions[targetDelete]
		if found && suggestion.Score > currentSuggestion.Score {
			continue
		}
		suggestions[targetDelete] = suggestion
	}
	// Precomputed words matches with precomputed dictionary words
	for _, targetDelete := range targetDeletes {
		deletes, found := s.deletes[targetDelete]
		if !found {
			continue
		}
		for _, delete := range deletes {
			suggestion := s.newSuggestion(targetDelete, delete, entry.Data, s.words[delete])
			currentSuggestion, found := suggestions[delete]
			if found && suggestion.Score > currentSuggestion.Score {
				continue
			}
			suggestions[delete] = suggestion
		}
	}
	// Turn map into ordered list of suggestions
	list := make([]Suggestion, 0, len(suggestions))
	for _, suggestion := range suggestions {
		list = append(list, suggestion)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Score < list[j].Score })
	return list
}

func (s *SymDel) newSuggestion(word, suggestion string, wordData, suggestionData *WordData) Suggestion {
	wordEntry := Entry{Word: word, Data: wordData}
	suggestionEntry := Entry{Word: suggestion, Data: suggestionData}
	return Suggestion{
		Entry: suggestionEntry,
		Score: s.distance(wordEntry, suggestionEntry),
	}
}

func CreateDeletes(w string, depth int) []string {
	word := []rune(w)
	n := len(word)
	depth = min(depth, n-1)
	deletes := make([]string, binomialsum(n-depth, n))
	store := func(res []rune) {
		deletes = append(deletes, string(res))
	}
	for i := n - depth; i < n; i++ {
		combinations(word, i, store)
	}
	return deletes
}

func combinations(word []rune, k int, f func([]rune)) {
	res := make([]rune, k)
	d := len(word) - k
	var comb func(int, int)
	comb = func(at, depth int) {
		if depth == k {
			f(res)
			return
		}
		for i := at; i <= d+depth; i++ {
			res[depth] = word[i]
			comb(i+1, depth+1)
		}
	}
	comb(0, 0)
}
