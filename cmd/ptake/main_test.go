package main

func Example_while() {
	goMain([]string{"ptake", "--while", "$PLINE != a3", "../../testdata/test1.txt"})
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

func Example_printHelp() {
	goMain([]string{"ptake", "--help"})
	// Output:
	// ptake version 1.0.0 (tamada/peripherals 1.0.0)
	// ptake [OPTIONS] [FILEs...]
	// OPTIONS
	//     -b, --bytes <NUMBER>       take NUMBER bytes (same as head command).
	//     -n, --lines <NUMBER>       take NUMBER lines (same as head command).
	//     -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
	//     -w, --while <PREDICATE>    take lines while PREDICATE is satisfied.
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
