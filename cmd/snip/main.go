package main

import (
	"errors"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/peripherals"
	"github.com/tamada/peripherals/cmd/common"
	"github.com/tamada/peripherals/errs"
	"github.com/tamada/peripherals/snip"
)

type SnipOpts struct {
	Number int
	snip.Snipper
	common.HeaderOpts
	common.HelpOptions
}

func (opts *SnipOpts) isInvalid() bool {
	return opts.Head < 0 && opts.Tail < 0 && opts.Number < 0
}

func (opts *SnipOpts) isOnlyNumberSet() bool {
	return opts.Head < 0 && opts.Tail < 0 && opts.Number > 0
}

func (opts *SnipOpts) Validate() error {
	if opts.isInvalid() {
		return errors.New("no lines specified. either options of -H, -T, -N must be specified")
	}
	if opts.isOnlyNumberSet() {
		opts.Head = opts.Number
		opts.Tail = opts.Number
	}
	return nil
}

func helpMessage(appName string, fs *flag.FlagSet) string {
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...]
OPTIONS
%s
FILE
  gives file name for the input. if this argument is single dash ('-') or absent,
  it reads strings from STDIN.
  if more than a single file is specified, each file is separated by a header
  consisting of the string '==> XXX <==' where 'XXX' is the name of the file.`, appName, fs.FlagUsages())
}

func buildFlags() (*flag.FlagSet, *SnipOpts) {
	var opts = &SnipOpts{}
	var flags = flag.NewFlagSet("pskip", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("pskip", flags)) }
	flags.IntVarP(&opts.Head, "head", "H", -1, "print first HEAD lines (same as head command).")
	flags.IntVarP(&opts.Tail, "tail", "T", -1, "print last TAIL lines (same as tail command).")
	flags.IntVarP(&opts.Number, "number", "N", 5, "print first and last lines (default is 5).")
	flags.BoolVarP(&opts.LineNumber, "line-number", "n", false, "print line number with output lines.")
	flags.BoolVarP(&opts.SkipSnip, "no-snip-sign", "s", false, "suppress printing of snip sign and the number of snipped lines.")
	flags.BoolVarP(&opts.NoHeader, "no-header", "q", false, "suppress printing of headers when multiple files are being examined.")
	flags.BoolVarP(&opts.HelpFlag, "help", "h", false, "print this message and exit")
	flags.BoolVarP(&opts.VersionFlag, "version", "v", false, "print the version information and exit")
	flags.SortFlags = false
	return flags, opts
}

func perform(opts *SnipOpts, args []string) error {
	if err := opts.Validate(); err != nil {
		return err
	}
	center := errs.New()
	for _, arg := range args {
		if opts.IsPrintHeader(args) {
			fmt.Printf("===> %s <===\n", arg)
		}
		center.Push(opts.PerformEach(arg, os.Stdout))
	}
	if len(args) == 0 {
		center.Push(opts.PerformEach("", os.Stdout))
	}
	return center.SelfOrNil()
}

func printHelp(opts *SnipOpts, fs *flag.FlagSet) int {
	if opts.VersionFlag {
		fmt.Println(peripherals.Version("snip"))
	}
	if opts.HelpFlag {
		fmt.Println(helpMessage("snip", fs))
	}
	return 0
}

func goMain(args []string) int {
	flags, opts := buildFlags()
	if err := flags.Parse(args); err != nil {
		return common.PrintError(err, 1)
	}
	if opts.IsHelp() {
		return printHelp(opts, flags)
	}
	return common.PrintError(perform(opts, flags.Args()[1:]), 2)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
