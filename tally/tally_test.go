package tally

import (
	"reflect"
	"testing"
)

func TestWordsCountsCaseInsensitively(t *testing.T) {
	got := Words("Go go GO lab")
	if got["go"] != 3 || got["lab"] != 1 {
		t.Fatalf("Words = %v", got)
	}
}

func TestTop(t *testing.T) {
	fruit := map[string]int{"pear": 3, "apple": 2, "banana": 1, "cherry": 1}
	tests := []struct {
		name   string
		counts map[string]int
		n      int
		want   []string
	}{
		{"descending counts", map[string]int{"a": 1, "b": 3, "c": 2}, 3, []string{"b", "c", "a"}},
		{"ties alphabetical at cutoff", fruit, 3, []string{"pear", "apple", "banana"}},
		{"all ties alphabetical", map[string]int{"z": 2, "m": 2, "a": 2}, 2, []string{"a", "m"}},
		{"n one", fruit, 1, []string{"pear"}},
		{"n equals distinct", fruit, 4, []string{"pear", "apple", "banana", "cherry"}},
		{"n exceeds distinct", fruit, 10, []string{"pear", "apple", "banana", "cherry"}},
		{"empty counts", map[string]int{}, 3, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Top(tt.counts, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Top(%v, %d) = %v, want %v", tt.counts, tt.n, got, tt.want)
			}
		})
	}
}
