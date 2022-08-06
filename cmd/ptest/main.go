package main

import (
	"fmt"
	"github.com/tamada/peripherals"
	"github.com/tamada/peripherals/ptest"
	"os"
)

const VERSION = "0.9.1"

func helpMessage(prog string) string {
	return fmt.Sprintf(`%s <expression>
file operations
    -b|-c|-d|-e|-f|-g|-k|-p|-r|-s|-u|-w|-x|-L|-O|-G|-S file
    -t file_descriptor
    file1 -nt|-ot|-ef file2
string operations
    [-n|-z] string
    s1 =|==|!=|>|>=|<|<=|-starts|-ends|-contains s2
number operations
    n1 -eq|-ne|-gt|-ge|-lt|-le n2
combination operations
    ! expression
    expression -a expression
    expression -o expression
    ( expression )
other operations
    --help               print this message and exit.
    --version            print the version information and exit.`, prog)
}

func perform(tokens []string) (bool, error) {
	return ptest.New(tokens).Eval()
}

func printHelp(appName string, versionFlag, helpFlag bool) int {
	if versionFlag {
		fmt.Println(peripherals.Version("ptest"))
	}
	if helpFlag {
		fmt.Println(helpMessage("ptest"))
	}
	return 0
}

func findHelpFlag(args []string) (bool, bool) {
	versionFlag := false
	helpFlag := false
	for _, arg := range args {
		if arg == "--version" {
			versionFlag = true
		} else if arg == "--help" {
			helpFlag = true
		}
	}
	return helpFlag, versionFlag
}

func goMain(args []string) int {
	helpFlag, versionFlag := findHelpFlag(args[1:])
	if helpFlag || versionFlag {
		return printHelp("ptest", versionFlag, helpFlag)
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
