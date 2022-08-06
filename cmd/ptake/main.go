package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/tamada/peripherals"
	"github.com/tamada/peripherals/cmd/common"
	"github.com/tamada/peripherals/errs"
	"github.com/tamada/peripherals/ptake"
	"os"
)

func helpMessage(appName string) string {
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...]
OPTIONS
    -b, --bytes <NUMBER>       take NUMBER bytes (same as head command).
    -n, --lines <NUMBER>       take NUMBER lines (same as head command).
    -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
    -w, --while <PREDICATE>    take lines while PREDICATE is satisfied.
                               we can use the variable PLINE and PLINECOUNT
                               which are the current line and its line number in the PREDICATE.
    -q, --no-header            suppress printing of headers when multiple files are being examined.

    -h, --help                 print this message and exit.
    -v, --version              print the version information and exit.
FILE
    gives file name for the input. if this argument is single dash ('-') or absent,
    it reads strings from STDIN.
    if more than a single file is specified, each file is separated by a header
    consisting of the string '==> XXX <==' where 'XXX' is the name of the file.`, appName)
}

func buildFlags() (*flag.FlagSet, *common.Options) {
	var opts = common.New()
	var flags = flag.NewFlagSet("ptake", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("ptake")) }
	flags.BoolVarP(&opts.NoHeader, "no-header", "q", false, "suppress printing of headers when multiple files are being examined.")
	flags.StringVarP(&opts.Keyword, "until", "u", "", "take lines until KEYWORD is appeared.")
	flags.IntVarP(&opts.Bytes, "bytes", "b", -1, "take BYTES bytes (same as head command).")
	flags.IntVarP(&opts.Lines, "lines", "n", -1, "take NUMBER lines (same as head command).")
	flags.BoolVarP(&opts.HelpFlag, "help", "h", false, "print this message and exit")
	flags.StringVarP(&opts.Predicate, "while", "w", "", "take lines while PREDICATE is satisfied.")
	return flags, opts
}

func perform(opts *common.Options, args []string) error {
	if err := opts.Validate(); err != nil {
		return err
	}
	out := bufio.NewWriter(os.Stdout)
	taker, err := ptake.New(opts.TType, opts)
	if err != nil {
		return err
	}
	center := errs.New()
	for _, arg := range args {
		if opts.IsPrintHeader(args) {
			fmt.Printf("===> %s <===\n", arg)
		}
		err := ptake.PerformEach(taker, arg, out)
		center.Push(err)
	}
	if len(args) == 0 {
		err := ptake.PerformEach(taker, "", out)
		center.Push(err)
	}
	return center.SelfOrNil()
}

func printHelp(opts *common.Options, appName string) int {
	if opts.VersionFlag {
		fmt.Println(peripherals.Version(appName))
	}
	if opts.HelpFlag {
		fmt.Println(helpMessage(appName))
	}
	return 0
}

func goMain(args []string) int {
	flags, opts := buildFlags()
	if err := flags.Parse(args); err != nil {
		return common.PrintError(err, 1)
	}
	if opts.IsHelp() {
		return printHelp(opts, "ptake")
	}
	return common.PrintError(perform(opts, flags.Args()[1:]), 2)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
