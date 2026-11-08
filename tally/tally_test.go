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

func TestWordsPunctuation(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "trim ascii boundary punctuation and merge case",
			text: "Go, go! GO.",
			want: map[string]int{"go": 3},
		},
		{
			name: "preserve internal apostrophes and hyphens",
			text: "don't well-known",
			want: map[string]int{"don't": 1, "well-known": 1},
		},
		{
			name: "trim unicode boundary punctuation",
			text: "«Go» “GO”",
			want: map[string]int{"go": 2},
		},
		{
			name: "punctuation-only tokens omitted",
			text: "... — !!!",
			want: map[string]int{},
		},
		{
			name: "empty input",
			text: "",
			want: map[string]int{},
		},
		{
			name: "whitespace-only input",
			text: "   \t\n  ",
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Words(tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Words(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}
