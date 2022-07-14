# pskip

## Usage

```sh
pskip [OPTIONS] [FILEs...]
OPTIONS
    -b, --bytes <BYTES>        skip BYTES bytes.
    -n, --lines <NUMBER>       skip NUMBER lines.
    -u, --until <KEYWORD>      skip lines until KEYWORD is appeared.
    -w, --while <PREDICATE>    skip lines while PREDICATE is satisfied.
FILEs
    if given file is "-" or absent, it reads from STDIN.
    if more than a single file is specified, each file is separated by a header 
    consisting of the string "==> XXX <==" where "XXX" is the name of the file.
```

