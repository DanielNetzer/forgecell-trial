package tally

import "testing"

func TestWordsCountsCaseInsensitively(t *testing.T) {
	got := Words("Go go GO lab")
	if got["go"] != 3 || got["lab"] != 1 {
		t.Fatalf("Words = %v", got)
	}
}
