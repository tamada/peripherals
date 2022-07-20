# peripheral utillity commands

utility commands for the shell by contrasting [GNU coreutils](https://www.gnu.org/software/coreutils/).

## How to install

```sh
brew tap tamada/brew
brew install peripherals
```

## implemented commands

* [`puniq`](#puniq)
* [`ptake`](#ptake)
* [`pskip`](#pskip)

### `puniq`

```sh
puniq [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
    -h, --help            print this message.

INPUT       gives file name of input.  If argument is single dash ('-')
            or absent, the program read strings from stdin.
OUTPUT      represents the destination.
```

### `ptake`

```sh
ptake [OPTIONS] [FILEs...]
OPTIONS
    -b, --bytes <NUMBER>       take NUMBER bytes (same as head command).
    -n, --lines <NUMBER>       take NUMBER lines (same as head command).
    -u, --until <KEYWORD>      take lines until KEYWORD is appeared.
    -w, --while <PREDICATE>    take lines while PREDICATE is satisfied. 
                               we can use the variable CLINE which is the current line
                               in the PREDICATE.
                               this option is unavailable on Windows platform.
    -q, --no-header            suppress printing of headers when multiple files are being examined.

    -h, --help                 print this message and exit.
    -v, --version              print version and exit.
FILE
    gives file name for the input. if this argument is single dash ("-") or absent,
    it reads strings from STDIN.
    if more than a single file is specified, each file is separated by a header 
    consisting of the string "==> XXX <==" where "XXX" is the name of the file.
```

### `pskip`


## License

* MIT License
