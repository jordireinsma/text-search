package main

import (
	"fmt"

	"github.com/jordi-reinsma/text-search/symdel"
)

func main() {
	s := symdel.New()
	s.Insert(symdel.Entry{
		Word: "jack",
		Data: &symdel.WordData{"num": 69, "nice": "yes"}},
	)
	s.Insert(symdel.Entry{
		Word: "bob",
		Data: &symdel.WordData{"count": 44, "nice": false}},
	)

	fmt.Println(s.Find("jack"))
	fmt.Println(s.Find("bob"))

	s.Insert(symdel.Entry{Word: "jacko"})
	s.Insert(symdel.Entry{Word: "backs"})
	s.Insert(symdel.Entry{Word: "jock"})

	fmt.Println(s.Find("jackoss"))

	suggs := s.Suggest(symdel.Entry{Word: "jack"})
	for _, sugg := range suggs {
		fmt.Println(sugg.Entry.Word, sugg.Score)
	}

	suggs = s.Suggest(symdel.Entry{Word: "back"})
	for _, sugg := range suggs {
		fmt.Println(sugg.Entry.Word, sugg.Score)
	}
}
