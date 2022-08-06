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
	// puniq [OPTIONS] [INPUT [OUTPUT]]
	// OPTIONS
	//     -a, --adjacent        delete only adjacent duplicated lines.
	//     -d, --delete-lines    only prints deleted lines.
	//     -i, --ignore-case     case sensitive.
	//
	//     -h, --help            print this message and exit.
	//     -v, --version         print the version information and exit.
	// INPUT                     gives file name of input.  If argument is single dash ('-')
	//                           or absent, the program read strings from stdin.
	// OUTPUT                    represents the destination.
}
