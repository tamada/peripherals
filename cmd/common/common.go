package common

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
	"github.com/tamada/peripherals/ptake"
)

type HelpOptions struct {
	HelpFlag    bool
	VersionFlag bool
}

func (opts *HelpOptions) IsHelp() bool {
	return opts.HelpFlag || opts.VersionFlag
}

type HeaderOpts struct {
	NoHeader bool
}

func (opts *HeaderOpts) IsPrintHeader(args []string) bool {
	return !opts.NoHeader && len(args) > 1
}

type Options struct {
	Lines     int
	Bytes     int
	Keyword   string
	Predicate string
	TType     ptake.TakerType
	HeaderOpts
	HelpOptions
}

func New() *Options {
	return &Options{Lines: -1, Bytes: -1, Keyword: "", Predicate: ""}
}

func (opts *Options) Int() int {
	if opts.TType == ptake.LINE {
		return opts.Lines
	} else if opts.TType == ptake.BYTES {
		return opts.Bytes
	}
	return -1
}

func (opts *Options) String() string {
	if opts.TType == ptake.WHILE {
		return opts.Predicate
	} else if opts.TType == ptake.UNTIL {
		return opts.Keyword
	}
	return ""
}

func (opts *Options) Validate() error {
	if opts.HelpFlag {
		return nil
	}
	bs := bitset.New(4)
	if opts.Bytes >= 0 {
		bs.Set(1)
		opts.TType = ptake.BYTES
	}
	if opts.Lines >= 0 {
		bs.Set(10)
		opts.TType = ptake.LINE
	}
	if opts.Keyword != "" {
		bs.Set(100)
		opts.TType = ptake.UNTIL
	}
	if opts.Predicate != "" {
		bs.Set(1000)
		opts.TType = ptake.WHILE
	}
	if !bs.Any() {
		return fmt.Errorf("should specify at least one criterion")
	}
	if bs.Count() != 1 {
		return fmt.Errorf("multiple criteria should not accept")
	}
	return nil
}

func PrintError(err error, statusOnError int) int {
	if err == nil {
		return 0
	}
	fmt.Println(err.Error())
	return statusOnError
}
