package main

func Example_uniq2() {
	goMain([]string{"puniq", "-i", "../../testdata/test1.txt"})
	// Output:
	// a1
	// a2
	// a3
	// a4
}

func Example_printHelp() {
	goMain([]string{"puniq", "--help"})
	// Output:
	// puniq version 2.0.0 (tamada/peripherals 0.9.0)
	// puniq [OPTIONS] [INPUT [OUTPUT]]
	// OPTIONS
	//     -a, --adjacent        delete only adjacent duplicated lines.
	//     -d, --delete-lines    only prints deleted lines.
	//     -i, --ignore-case     case sensitive.
	//
	//     -h, --help            print this message.
	// INPUT                     gives file name of input.  If argument is single dash ('-')
	//                           or absent, the program read strings from stdin.
	// OUTPUT                    represents the destination.
}
