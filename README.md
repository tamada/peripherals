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
    -b, --bytes <BYTES>        take BYTES bytes (same as head command).
    -n, --lines <NUMBER>       take NUMBER lines (same as head command).
    -u, --until <KEYWORD>      take lines while KEYWORD is appeared.
    -w, --while <PREDICATE>    take lines until PREDICATE is satisfied. 
                               we can use the variable which CLINE shows the current line
                               in the PREDICATE.
    -q, --no-header            suppress printing of headers when multiple files are being examined.

FILE    if given file is "-" or absent, it reads from STDIN.
        if more than a single file is specified, each file is separated by a header 
        consisting of the string "==> XXX <==" where "XXX" is the name of the file.
```

### `pskip`


## License

* MIT License
