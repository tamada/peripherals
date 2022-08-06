package main

import (
	"fmt"
	"github.com/tamada/peripherals"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/peripherals/puniq"
)

func helpMessage(appName string) string {
	return fmt.Sprintf(`%s [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.

    -h, --help            print this message and exit.
    -v, --version         print the version information and exit.
INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.`, appName)
}

func printError(err error, statusCode int) int {
	if err == nil {
		return 0
	}
	fmt.Println(err.Error())
	return statusCode
}

func perform(flags *flag.FlagSet, opts *puniq.Parameters) int {
	var args, err = puniq.NewArguments(flags.Args()[1:])
	if err != nil {
		return printError(err, 1)
	}
	defer args.Close()
	err = args.Perform(opts)
	return printError(err, 2)
}

func printHelp(opts *options, appName string) int {
	if opts.versionFlag {
		fmt.Println(peripherals.Version(appName))
	}
	if opts.helpFlag {
		fmt.Println(helpMessage(appName))
	}
	return 0
}

func goMain(args []string) int {
	var flags, opts = buildFlagSet()
	var err = flags.Parse(args)
	if err != nil {
		return printError(err, 1)
	}
	if opts.helpFlag || opts.versionFlag {
		return printHelp(opts, "puniq")
	}
	return perform(flags, opts.params)
}

type options struct {
	params      *puniq.Parameters
	helpFlag    bool
	versionFlag bool
}

func buildFlagSet() (*flag.FlagSet, *options) {
	var opts = options{params: &puniq.Parameters{}}
	var flags = flag.NewFlagSet("uniq2", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage("puniq")) }
	flags.BoolVarP(&opts.params.Adjacent, "adjacent", "a", false, "delete only the adjacent duplicate lines")
	flags.BoolVarP(&opts.params.DeleteLines, "delete-lines", "d", false, "only prints deleted lines")
	flags.BoolVarP(&opts.params.IgnoreCase, "ignore-case", "i", false, "case sensitive")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message and exit")
	flags.BoolVarP(&opts.versionFlag, "version", "v", false, "print version information and exit")
	return flags, &opts
}

func main() {
	var exitStatus = goMain(os.Args)
	os.Exit(exitStatus)
}
