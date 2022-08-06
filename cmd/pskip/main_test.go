package main

func Example_while() {
	goMain([]string{"pskip", "--while", "$PLINE != a3", "../../testdata/test1.txt"})
	// Output:
	// a3
	// a4
	// a1
	// A1
}

func Example_until() {
	goMain([]string{"pskip", "--until", "a4", "../../testdata/test1.txt"})
	// Output:
	// a4
	// a1
	// A1
}

func Example_printHelp() {
	goMain([]string{"pskip", "--help"})
	// Output:
	// pskip version 1.0.0 (tamada/peripherals 0.9.0)
	// pskip [OPTIONS] [FILEs...]
	// OPTIONS
	//     -b, --bytes <NUMBER>       skip NUMBER bytes (same as head command).
	//     -n, --lines <NUMBER>       skip NUMBER lines (same as head command).
	//     -u, --until <KEYWORD>      skip lines until KEYWORD is appeared.
	//     -w, --while <PREDICATE>    skip lines while PREDICATE is satisfied.
	//                                we can use the variable PLINE and PLINECOUNT
	//                                which are the current line and its line number in the PREDICATE.
	//     -q, --no-header            suppress printing of headers when multiple files are being examined.
	//
	//     -h, --help                 print this message and exit.
	// FILE
	//     gives file name for the input. if this argument is single dash ('-') or absent,
	//     it reads strings from STDIN.
	//     if more than a single file is specified, each file is separated by a header
	//     consisting of the string '==> XXX <==' where 'XXX' is the name of the file.
}
