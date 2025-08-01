package symdel

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, fail bool) {
	if fail {
		t.Errorf("assertion failed")
	}
}

func TestSymDel_Find(t *testing.T) {
	s := New()
	s.Insert(Entry{Word: "jack"})
	s.Insert(Entry{Word: "bob"})

	fmt.Println(s.Find("jack"))
	fmt.Println(s.Find("bob"))
}
