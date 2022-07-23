package ptest

import (
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	os.Setenv("VALUE", "10")
	os.Setenv("DEVELOPER", "tamada")
	testdata := []struct {
		giveString []string
		wontError  bool
		wontFlag   bool
	}{
		{[]string{"hoge != fuga"}, false, true},
		{[]string{"-r ptest.go"}, false, true},
		{[]string{"1 -eq 1"}, false, true},
		{[]string{"$DEVELOPER", "==", "tamada"}, false, true},
		{[]string{"$VALUE -ne 9"}, false, true},
		{[]string{"$VALUE -gt 8"}, false, true},
		{[]string{"$VALUE -gt hoge"}, true, false},
	}
	for _, td := range testdata {
		e := New(td.giveString)
		flag, err := e.Eval()
		if !td.wontError && err != nil {
			t.Errorf("%v: wont no error, but got error (%s)", td.giveString, err.Error())
		}
		if td.wontError && err == nil {
			t.Errorf("%v: wont error, but got no error", td.giveString)
		}
		if flag != td.wontFlag {
			t.Errorf("%v: result did not match, wont %v, got %v", td.giveString, td.wontFlag, !td.wontFlag)
		}
	}
}
