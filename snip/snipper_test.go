package snip

import (
	"bytes"
	"testing"
)

func TestPerformEach(t *testing.T) {
	s := &Snipper{Head: 1, Tail: 1, SkipSnip: false, LineNumber: false}
	out := bytes.NewBuffer([]byte{})
	if err := s.PerformEach("../testdata/test1.txt", out); err != nil {
		t.Errorf("error should not be raised, but %s", err)
	}
	got := out.String()
	wont := "a1\n        ... snip ... (6 lines)\nA1\n"
	if got != wont {
		t.Errorf("result did not match, wont %s, got %s", wont, got)
	}
}
