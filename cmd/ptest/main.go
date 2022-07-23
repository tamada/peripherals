package main

import (
	"fmt"
	"github.com/tamada/peripherals/ptest"
	"os"
)

func printHelp(prog string) string {
	return fmt.Sprintf(`%s <expression>
file operations
    -b|-c|-d|-e|-f|-g|-h|-k|-p|-r|-s|-u|-w|-x|-L|-O|-G|-S file
    -t file_descriptor
    file1 -nt|-ot|-ef file2
string operations
    [-n|-z] string
    s1 =|==|!=|>|>=|<|<= s2
number operations
    n1 -eq|-ne|-gt|-ge|-lt|-le n2
combination operations
    ! expression
    expression -a expression
    expression -o expression
    ( expression )
other operations
    --help               print this message.`, prog)
}

func perform(tokens []string) (bool, error) {
	return ptest.New(tokens).Eval()
}

func goMain(args []string) int {
	for _, str := range args[1:] {
		if str == "--help" {
			fmt.Println(printHelp("ptest"))
			return 0
		}
	}
	result, err := perform(args[1:])
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return resultToInt(result)
}

func resultToInt(flag bool) int {
	if flag {
		return 0
	}
	return 1
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
