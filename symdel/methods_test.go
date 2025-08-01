package symdel

import "testing"

func TestDemerauLevenshteinDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want float64
	}{
		{"abc", "abc", 0.0},
		{"bob", "bobo", 1.0},
		{"bob", "boo", 1.0},
		{"bob", "bo", 1.0},
		{"bob", "obo", 2.0},
		{"bo", "ob", 2.0},
		{"bob", "jack", 4.0},
		{"", "abc", 3.0},
		{"abc", "", 3.0},
		{"", "", 0.0},
		{"aço", "abo", 1.0},
		{"ẽaço", "fabo", 2.0},
		{"abcba", "abxba", 1.0},
		{"abcba", "axcba", 1.0},
	}
	for _, test := range tests {
		got := DemerauLevenshteinDistance(
			Entry{Word: test.a, Data: nil}, Entry{Word: test.b, Data: nil},
		)
		if got != test.want {
			t.Errorf("f(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}
