package main

func Example_withNumber() {
	goMain([]string{"snip", "--line-number", "--number", "2", "../../testdata/test1.txt"})
	// Output:
	//      1  a1
	//      2  a1
	//         ... snip ... (4 lines)
	//      7  a1
	//      8  A1
}

func Example_headTail() {
	goMain([]string{"snip", "--no-snip-sign", "--head", "1", "--tail", "2", "../../testdata/test1.txt"})
	// Output:
	// a1
	// a1
	// A1
}

func Example_ShortLengthFile() {
	goMain([]string{"snip", "../../testdata/eval_script.sh"}) // 4 lines
	// Output:
	// #! /bin/sh
	//
	// echo "PLINE: $PLINE"
	// exec test "$PLINE" = "a1"
}

func Example_printHelp() {
	goMain([]string{"snip", "--help"})
	// Output:
	// snip [OPTIONS] [FILEs...]
	// OPTIONS
	//   -H, --head int       print the specified leading lines (same as head command). (default 5)
	//   -T, --tail int       print the specified tailing lines (same as tail command). (default 5)
	//   -N, --number int     print the specified number of leading and tailing lines. If this option
	//                        value is positive, it is treated as if this value were specified for -H
	//                        and -T. (default -1)
	//   -n, --line-number    print line number with output lines.
	//   -s, --no-snip-sign   suppress printing of snip sign and the number of snipped lines.
	//   -q, --no-header      suppress printing of headers when multiple files are being examined.
	//   -h, --help           print this message and exit.
	//   -v, --version        print the version information and exit.
	//
	// FILE
	//   gives file name for the input. if this argument is single dash ('-') or absent,
	//   it reads strings from STDIN.
	//   if more than a single file is specified, each file is separated by a header
	//   consisting of the string '==> XXX <==' where 'XXX' is the name of the file.
}
