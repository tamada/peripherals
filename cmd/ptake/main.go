package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bits-and-blooms/bitset"
	flag "github.com/spf13/pflag"
	"github.com/tamada/peripherals/errs"
	"github.com/tamada/peripherals/ptake"
)

const VERSION = "1.0.0"

type options struct {
	lines     int
	bytes     int
	keyword   string
	predicate string
	noHeader  bool
	helpFlag  bool
	ttype     ptake.TakerType
}

func newOptions() *options {
	return &options{lines: -1, bytes: -1, keyword: "", predicate: "", noHeader: false}
}

func (opts *options) Int() int {
	if opts.ttype == ptake.LINE {
		return opts.lines
	} else if opts.ttype == ptake.BYTES {
		return opts.bytes
	}
	return -1
}

func (opts *options) String() string {
	if opts.ttype == ptake.WHILE {
		return opts.predicate
	} else if opts.ttype == ptake.UNTIL {
		return opts.keyword
	}
	return ""
}

func (opts *options) validate() error {
	if opts.helpFlag {
		return nil
	}
	bs := bitset.New(4)
	if opts.bytes >= 0 {
		bs.Set(1)
		opts.ttype = ptake.BYTES
	}
	if opts.lines >= 0 {
		bs.Set(10)
		opts.ttype = ptake.LINE
	}
	if opts.keyword != "" {
		bs.Set(100)
		opts.ttype = ptake.UNTIL
	}
	if opts.predicate != "" {
		bs.Set(1000)
		opts.ttype = ptake.WHILE
	}
	if !bs.Any() {
		return fmt.Errorf("should specify at least one criterion")
	}
	if bs.Count() != 1 {
		return fmt.Errorf("multiple criteria should not accept")
	}
	return nil
}

func helpMessage(prog string) string {
	return fmt.Sprintf(`%s [OPTIONS] [FILEs...]
OPTIONS
    -b, --bytes <BYTES>        take BYTES bytes (same as head command).
    -n, --lines <NUMBER>       take NUMBER lines (same as head command).
    -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
    -w, --while <PREDICATE>    take lines while PREDICATE is satisfied. 
                               we can use the variable which CLINE shows 
                               the current line in the PREDICATE.
    -q, --no-header            suppress printing of headers when multiple
                               files are being examined.
FILE
    gives file name for the input. if this argument is single dash ("-") or absent,
    it reads strings from STDIN.
    if more than a single file is specified, each file is separated by a header 
    consisting of the string "==> XXX <==" where "XXX" is the name of the file.`, prog)
}

func buildFlags() (*flag.FlagSet, *options) {
	var opts = newOptions()
	var flags = flag.NewFlagSet("ptake", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("ptake")) }
	flags.BoolVarP(&opts.noHeader, "no-header", "q", false, "suppress printing of headers when multiple files are being examined.")
	flags.StringVarP(&opts.keyword, "until", "u", "", "take lines until KEYWORD is appeared.")
	flags.StringVarP(&opts.predicate, "while", "w", "", "take lines while PREDICATE is satisfied.")
	flags.IntVarP(&opts.bytes, "bytes", "b", -1, "takes BYTES bytes (same as head command).")
	flags.IntVarP(&opts.lines, "lines", "n", -1, "takes NUMBER lines (same as head command).")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message.")
	return flags, opts
}

func perform(opts *options, args []string) error {
	if err := opts.validate(); err != nil {
		return err
	}
	out := bufio.NewWriter(os.Stdout)
	taker, err := ptake.NewTaker(opts.ttype, opts)
	if err != nil {
		return err
	}
	center := errs.New()
	for _, arg := range args {
		err := ptake.PerformEach(taker, arg, out)
		center.Push(err)
	}
	return center.SelfOrNil()
}

func printError(err error, statusOnError int) int {
	if err == nil {
		return 0
	}
	fmt.Println(err.Error())
	return statusOnError
}

func goMain(args []string) int {
	flags, opts := buildFlags()
	if err := flags.Parse(args); err != nil {
		return printError(err, 1)
	}
	if opts.helpFlag {
		return printError(fmt.Errorf(helpMessage("ptake")), 0)
	}
	return printError(perform(opts, flags.Args()[1:]), 2)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
