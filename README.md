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
                               we can use the variable PTAKE_LINE and PTAKE_LINECOUNT
                               which are the current line and its number in the PREDICATE.
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


### `ptest`

```shell
ptest <expression>
file operation
    -b file      
    -c file
    -d file
    -e file
    -f file
    -g file
    -h file
    -k file
    -p file
    -r file
    -s file
    -t file_descriptor 
    -u file            true if file exists and its set user id flag is set.
    -w file            true if file exists and is writable.
    -x file            true if file exists and is executable.  If file is directory,
                       true indicates that file can be searched.
    -L file            true if file exists and its a symbolic link.
    -O file            true if file exists and its owner matches the effective user id of this process.
    -G file            true if file exists and its group matches the effective group id of this process.
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
    s1 < s2
    s1 <= s2
    s1 > s2
    s1 >= s2
numeric operation
    n1 -eq n2
    n1 -ne n2
    n1 -gt n2
    n1 -ge n2
    n1 -lt n2
    n1 -le n2
combined operation
    ! expression
    expression -a expression
    expression -o expression
    ( expression )
    
```


## License

* MIT License
