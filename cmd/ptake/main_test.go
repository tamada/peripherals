package main

import "testing"

// func Example_printHelp() {
// 	goMain([]string{"ptake", "--help"})
// 	// Output:
// 	// ptake [OPTIONS] [FILEs...]
// 	// OPTIONS
// 	//     -b, --bytes <BYTES>        take BYTES bytes (same as head command).
// 	//     -n, --lines <NUMBER>       take NUMBER lines (same as head command).
// 	//     -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
// 	//     -w, --while <PREDICATE>    take lines while PREDICATE is satisfied.
// 	//                                we can use the variable which CLINE shows
// 	//                                the current line in the PREDICATE.
// 	//     -q, --no-header            suppress printing of headers when multiple
// 	//                                files are being examined.
// 	// FILE
// 	//     gives file name for the input. if this argument is single dash ("-") or absent,
// 	//     it reads strings from STDIN.
// 	//     if more than a single file is specified, each file is separated by a header
// 	//     consisting of the string "==> XXX <==" where "XXX" is the name of the file.
// }

func TestValidate(t *testing.T) {
	opts1 := options{lines: 0, bytes: 0, keyword: "", predicate: "", noHeader: false}
	if err := opts1.validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
}
