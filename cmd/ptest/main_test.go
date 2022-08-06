package main

import (
	"testing"
)

func Example_printHelp() {
	goMain([]string{"ptest", "--help"})
	// Output:
	// ptest version 1.0.0 (tamada/peripherals 1.0.0)
	// ptest <expression>
	// file operations
	//     -b|-c|-d|-e|-f|-g|-k|-p|-r|-s|-u|-w|-x|-L|-O|-G|-S file
	//     -t file_descriptor
	//     file1 -nt|-ot|-ef file2
	// string operations
	//     [-n|-z] string
	//     s1 =|==|!=|>|>=|<|<=|-starts|-ends|-contains s2
	// number operations
	//     n1 -eq|-ne|-gt|-ge|-lt|-le n2
	// combination operations
	//     ! expression
	//     expression -a expression
	//     expression -o expression
	//     ( expression )
	// other operations
	//     --help               print this message.
}

func TestPTest(t *testing.T) {
	testdata := []struct {
		args       []string
		wontResult int
	}{
		{[]string{"-f", "main.go"}, 0},
	}
	for _, td := range testdata {
		gotResult := goMain(td.args)
		if gotResult != td.wontResult {
			t.Errorf("%v: result did not match wont %d, got %d", td.args, td.wontResult, gotResult)
		}
	}
}
