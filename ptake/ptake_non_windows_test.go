//go:build !windows

package ptake

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPerformEachForNonWindows(t *testing.T) {
	testdata := []struct {
		label    string
		taker    Taker
		file     string
		wontData string
	}{
		{"while2", createTaker(WHILE, newSValue("../testdata/eval_script.sh")), "../testdata/test1.txt", "a1\na1\n"},
	}

	for _, td := range testdata {
		buffer := bytes.NewBuffer([]byte{})
		err := PerformEach(td.taker, td.file, bufio.NewWriter(buffer))
		gotData := buffer.String()
		if err != nil {
			t.Errorf("unknown error: %s", err.Error())
		}
		if gotData != td.wontData {
			t.Errorf("%s: result did not match, wont %s, got %s", td.label, td.wontData, gotData)
		}
	}
}
