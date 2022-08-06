# peripheral utility commands

[![License](https://img.shields.io/badge/License-MIT-green)](https://github.com/tamada/peripherals/blob/main/LICENSE)
[![Version](https://img.shields.io/badge/Version-v${VERSION}-green)](https://github.com/tamada/btmeister/releases/tag/v${VERSION})

peripheral utility commands for the shell by contrasting [GNU coreutils](https://www.gnu.org/software/coreutils/).

## How to install

```sh
brew tap tamada/brew
brew install peripherals
```

## implemented commands

* [`puniq`](#puniq)
* [`ptake`](#ptake)
* [`pskip`](#pskip)
* [`ptest`](#ptest)

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
                               we can use the variable PLINE and PLINECOUNT
                               which are the current line and its line number in the PREDICATE.
    -q, --no-header            suppress printing of headers when multiple files are being examined.

    -h, --help                 print this message and exit.
FILE
    gives file name for the input. if this argument is single dash ("-") or absent,
    it reads strings from STDIN.
    if more than a single file is specified, each file is separated by a header
    consisting of the string "==> XXX <==" where "XXX" is the name of the file.
```

### `pskip`

```shell
skip [OPTIONS] [FILEs...]
OPTIONS
    -b, --bytes <NUMBER>       skip NUMBER bytes (same as head command).
    -n, --lines <NUMBER>       skip NUMBER lines (same as head command).
    -u, --until <KEYWORD>      skip lines until KEYWORD is appeared.
    -w, --while <PREDICATE>    skip lines while PREDICATE is satisfied.
                               we can use the variable PLINE and PLINECOUNT
                               which are the current line and its line number in the PREDICATE.
    -q, --no-header            suppress printing of headers when multiple files are being examined.

    -h, --help                 print this message and exit.
FILE
    gives file name for the input. if this argument is single dash ("-") or absent,
    it reads strings from STDIN.
    if more than a single file is specified, each file is separated by a header
    consisting of the string "==> XXX <==" where "XXX" is the name of the file.
```

### `ptest`

```shell
ptest <expression>
file operation
    -b file            true if file exists and is a block special file.
    -c file            true if file exists and is a character special file.
    -d file            true if file exists and is a directory.
    -e file            true if file exists (regardless of type).
    -f file            true if file exists and is a regular file.
    -g file            true if file exists and its set group ID flag is set.
    -k file            true if file exists and its sticky bit is set.
    -p file            true if file is a named pipe (FIFO).
    -r file            true if file exists and is readable.
    -s file            true if file exists and has a size greater than zero.
    -t file_descriptor true if the file whose file descriptor number is file_descriptor is open and is
                       associated with a terminal.
    -u file            true if file exists and its set user id flag is set.
    -w file            true if file exists and is writable.
    -x file            true if file exists and is executable.  If file is directory,
                       true indicates that file can be searched.
    -L file            true if file exists and its a symbolic link.
    -O file            true if file exists and its owner matches the effective user id of this process (not support on Windows platform).
    -G file            true if file exists and its group matches the effective group id of this process (not support on Windows platform).
    -S file            true if file exists and its a socket.
    file1 -nt file2    true if file1 exists and is newer than file2.
    file1 -ot file2    true if file1 exists and is older than file2.
    file1 -ef file2    true if file1 and file2 exist and refer to the same file.
string operation
    -n string          true if the length of string is nonzero.
    -z string          true if the length of string is zero.
    string             true if string is not the null string.
    s1 = s2            true if the strings s1 and s2 are identical.
    s1 == s2           true if the strings s1 and s2 are identical, same as s1 = s2.
    s1 != s2           true if the strings s1 and s2 are not identical.
    s1 -starts s2      true if string s1 starts with s2.
    s1 -ends s2        true if string s1 ends with s2.
    s1 -contains s2    true if string s1 contains s2.
    s1 < s2            true if string s1 comes before s2 based on the binary value of their characters.
    s1 <= s2           true if string s1 comes before or equals to s2 based on the binary value of their characters.
    s1 > s2            true if string s1 comes after s2 based on the binary value of their characters.
    s1 >= s2           true if string s1 comes after or equals to s2 based on the binary value of their characters.
numeric operation
    n1 -eq n2          true if the integers n1 and n2 are algebraically equal.
    n1 -ne n2          true if the integers n1 and n2 are not algebraically equal.
    n1 -gt n2          true if the integers n1 is algebraically greater than the integer n2.
    n1 -ge n2          true if the integers n1 is algebraically greater than or equals to the integer n2.
    n1 -lt n2          true if the integers n1 is algebraically less than the integer n2.
    n1 -le n2          true if the integers n1 is algebraically less than or equals to the integer n2.
combined operation
    ! expression                true if the expression is false.
    expression -a expression    true if both expression1 and expression2 are true.  The -a operator has higher precedence than the -o operator.
    expression -o expression    true if either expression1 or expression2 are true.  The -o operator has lower precedence than the -a operator.
    ( expression )              true if expression is true
```

## License

* MIT License
