package ptake

import (
	"bufio"
	"bytes"
	"testing"
)

type input struct {
	iValue int
	sValue string
}

func newIValue(value int) *input {
	return &input{iValue: value}
}

func newSValue(value string) *input {
	return &input{sValue: value}
}

func (i *input) Int() int {
	return i.iValue
}

func (i *input) String() string {
	return i.sValue
}

func createTaker(t TakerType, data InputData) Taker {
	taker, _ := New(t, data)
	return taker
}

func TestPerformEach(t *testing.T) {
	testdata := []struct {
		label    string
		taker    Taker
		file     string
		wontData string
	}{
		{"bytes", createTaker(BYTES, newIValue(5)), "../testdata/test1.txt", "a1\na1"},
		{"line", createTaker(LINE, newIValue(3)), "../testdata/test1.txt", "a1\na1\na2\n"},
		{"until", createTaker(UNTIL, newSValue("a3")), "../testdata/test1.txt", "a1\na1\na2\na2\n"},
		{"while1", createTaker(WHILE, newSValue("../testdata/eval_script.sh")), "../testdata/test1.txt", "a1\na1\n"},
		{"while2", createTaker(WHILE, newSValue("$PLINE = a1")), "../testdata/test1.txt", "a1\na1\n"},
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
