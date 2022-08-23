---
title: ":runner: Usage"
---

* [`puniq`](#puniq)
    {{< label_button label="Description" href="#-description" >}} 
    {{< label_button label="Usage" href="#-usage" >}} 
    {{< label_button label="Demo" href="#-demo" >}} 
    {{< label_button label="Utility" href="#-utility" >}} 
* [`ptake`](#ptake)
    {{< label_button label="Description" href="#-description-1" >}} 
    {{< label_button label="Usage" href="#-usage-1" >}} 
* [`pskip`](#pskip)
    {{< label_button label="Description" href="#-description-2" >}} 
    {{< label_button label="Usage" href="#-usage-2" >}} 
* [`ptest`](#ptest)
    {{< label_button label="Description" href="#-description-2" >}} 
    {{< label_button label="Usage" href="#-usage-2" >}} 

## `puniq`

### :speaking_head: Description

Delete duplicated lines.

GNU core utilities have `uniq` command for deleting duplicate lines.
However, `uniq` command deletes only continuous duplicate lines.
When deleting not continuous duplicate lines, we use `sort` command together, in that case, the order of the list was not kept.

We want to delete not continuous duplicated lines with remaining the order.
This command is the same as [`uniq2`](https://github.com/tamada/uniq2)

### :runner: Usage

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

### :bicyclist: Demo

```sh
$ cat -n testdata/test1.txt
1	a1
2	a1  # <- is the duplicate of the previous line.
3	a2
4	a2  # <- is the duplicate of the previous line.
5	a3
6	a4
7	a1  # <- is the duplicate of the first line.
8	A1
$ puniq testdata/test1.txt
a1
a2
a3
a4
A1
$ puniq -a testdata/test1.txt  # same result as uniq command.
a1
a2
a3
a4
a1  # <- this line is not deleted.
A1
$ puniq -i testdata/test1.txt # ignore case
a1
a2
a3
a4
$ puniq -d testdata/test1.txt # print delete lines.
a1
a2
a1
```

#### :surfer: Utility

Delete duplicate entries in PATH

```sh
export PATH=$(echo $PATH | tr : '\n' | puniq | paste -s -d : -)
```

* `tr : '\n'` replaces `:` to `\n` of data from STDIN,
* `uniq2` deletes duplicate lines from the result of `tr`, and
* `paste -s -d : -` joins given strings with `:`.

## `ptake`

### :speaking_head: Description

take lines while satisfying the predicate.

### :runner: Usage

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

## `pskip`

### :speaking_head: Description

skip lines while satisfying the predicate.

### :runner: Usage

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

## `ptest`

### :speaking_head: Description

another implementation of `test` command for evaluating the predicate of `pskip` and `ptake`.

### :runner: Usage

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



