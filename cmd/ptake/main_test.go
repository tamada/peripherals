package main

import (
	"testing"
)

func Example_while() {
	goMain([]string{"ptake", "--while", "$CLINE != a3", "../../testdata/test1.txt"})
	// Output:
	// a1
	// a1
	// a2
	// a2
}

func Example_until() {
	goMain([]string{"ptake", "--until", "a4", "../../testdata/test1.txt"})
	// Output:
	// a1
	// a1
	// a2
	// a2
	// a3
}

//func Example_printHelp() {
//	goMain([]string{"ptake", "--help"})
//	// Output:
//	// ptake [OPTIONS] [FILEs...]
//	// OPTIONS
//	//     -b, --bytes <NUMBER>       take NUMBER bytes (same as head command).
//	//     -n, --lines <NUMBER>       take NUMBER lines (same as head command).
//	//     -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
//	//     -w, --while <PREDICATE>    take lines while PREDICATE is satisfied.
//	//                                we can use the variable PTAKE_LINE and PTAKE_LINECOUNT
//	//                                which are the current line and its number in the PREDICATE.
//	//     -q, --no-header            suppress printing of headers when multiple
//	//                                files are being examined.
//	//
//	//     -h, --help                 print this message and exit.
//	//     -v, --version              print version and exit.
//	// FILE
//	//     gives file name for the input. if this argument is single dash ("-") or absent,
//	//     it reads strings from STDIN.
//	//     if more than a single file is specified, each file is separated by a header
//	//     consisting of the string "==> XXX <==" where "XXX" is the name of the file.
//}

func TestValidate(t *testing.T) {
	opts1 := options{lines: 0, bytes: 0, keyword: "", predicate: "", noHeader: false}
	if err := opts1.validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
	opts2 := options{lines: 10, bytes: 10, keyword: "", predicate: "", noHeader: false}
	if err := opts2.validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
}
