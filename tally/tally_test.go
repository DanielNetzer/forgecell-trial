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

func TestWordsTrimsPunctuation(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{"trailing punctuation", "Go, go! GO.", map[string]int{"go": 3}},
		{"unicode punctuation", "«go» ¡go!", map[string]int{"go": 2}},
		{"inner apostrophe", "Don't don't", map[string]int{"don't": 2}},
		{"inner hyphen", "well-known (well-known),", map[string]int{"well-known": 2}},
		{"punctuation only", "-- !!! ... go", map[string]int{"go": 1}},
		{"only punctuation", "?! ...", map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Words(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Words(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}
