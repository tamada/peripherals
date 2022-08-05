package ptest

import (
	"fmt"
	"github.com/tamada/peripherals"
	"os"
	"strconv"
	"strings"
	"syscall"
)

type Expression struct {
	Original string
	tokens   []string
	index    int
	envs     []*envVariable
}

type envVariable struct {
	key   string
	value string
}

func findEnvs() []*envVariable {
	results := []*envVariable{}
	for _, env := range os.Environ() {
		items := strings.Split(env, "=")
		results = append(results, &envVariable{key: items[0], value: items[1]})
	}
	return results
}

func updateTokens(tokens []string) []string {
	results := []string{}
	for _, str := range tokens {
		items := strings.Split(str, " ")
		for _, item := range items {
			results = append(results, item)
		}
	}
	return results
}

func New(expression []string) *Expression {
	envs := findEnvs()
	tokens := updateTokens(expression)
	return &Expression{Original: strings.Join(expression, " "), tokens: tokens, index: 0, envs: envs}
}

func (e *Expression) SetEnv(key, value string) bool {
	if contains(e.envs, key) {
		for _, env := range e.envs {
			if env.key == key {
				env.value = value
			}
		}
	} else {
		e.envs = append(e.envs, &envVariable{key: key, value: value})
	}
	return true
}

func contains(variables []*envVariable, key string) bool {
	for _, env := range variables {
		if env.key == key {
			return true
		}
	}
	return false
}

func (e *Expression) current() (string, bool) {
	if len(e.tokens) > e.index {
		return e.tokens[e.index], true
	}
	return "", false
}

func (e *Expression) next() (string, bool) {
	e.index = e.index + 1
	return e.current()
}

func eval(e *Expression) (bool, error) {
	c, ok := e.current()
	if ok && c == "!" {
		e.next()
		result, err := e.Term()
		return !result, err
	} else {
		return e.Term()
	}
}

func (e *Expression) Eval() (bool, error) {
	r1, err := eval(e)
	if err != nil {
		return false, err
	}
	op, ok := e.current()
	if !ok {
		fmt.Errorf("%s: parse error", e.Original)
	}
	for op == "-a" || op == "-o" {
		_, ok := e.next()
		if !ok {
			return false, fmt.Errorf("%s: parse error", e.Original)
		}
		r2, err := e.Eval()
		if err != nil {
			return false, err
		}
		if op == "-a" {
			r1 = r1 && r2
		} else {
			r1 = r1 || r2
		}
	}
	return r1, nil
}

func operation(opts, target string) (bool, error) {
	stat, err := os.Stat(target)
	switch opts {
	case "-n":
		return len(target) > 0, nil
	case "-z":
		return len(target) == 0, nil

	case "-b": // block special file
		return err == nil && stat.Mode() == os.ModeDevice && stat.Mode() != os.ModeCharDevice, nil
	case "-c": // character special file
		return err == nil && stat.Mode() == os.ModeCharDevice, nil
	case "-d": // directory
		return err == nil && stat.IsDir(), nil
	case "-f": // regular file
		return err == nil && stat.Mode().IsRegular(), nil
	case "-g": // group ID flag
		return err == nil && stat.Mode() == os.ModeSetgid, nil
	case "-h", "-L": // symbolic link
		return err == nil && stat.Mode() == os.ModeSymlink, nil
	case "-k": // sticky bit set
		return err == nil && stat.Mode() == os.ModeSticky, nil
	case "-p": // named pipe
		return err == nil && stat.Mode() == os.ModeNamedPipe, nil
	case "-r": // readable
		return err == nil && stat.Mode()&(1<<(9-1)) != 0, nil
	case "-s": // not empty file
		return err == nil && stat.Size() > 0, nil
	case "-u": // set uid flag
		return err == nil && stat.Mode() == os.ModeSetuid, nil
	case "-w": // writable
		return err == nil && stat.Mode()&(1<<(9-2)) != 0, nil
	case "-x": // executable
		return err == nil && stat.Mode()&(1<<(9-3)) != 0, nil
	case "-S": // socket
		return err == nil && stat.Mode() == os.ModeSocket, nil
	case "-O": // owner matches the effective user of this process
		if peripherals.IsWindows() {
			return false, nil
		}
		info := stat.Sys().(*syscall.Stat_t)
		return err == nil && int(info.Uid) == os.Getuid(), nil
	case "-G": // owner matches the effective user of this process
		if peripherals.IsWindows() {
			return false, nil
		}
		info := stat.Sys().(*syscall.Stat_t)
		return err == nil && int(info.Gid) == os.Getgid(), nil
	}
	return true, nil
}

var stringOpts = []string{
	"==", "=", "!=", ">", ">=", "<", "<=",
}
var numberOpts = []string{
	"-eq", "-ne", "-lt", "-le", "-gt", "-ge",
}
var fileOpts = []string{
	"-nt", "-ot", "-ef",
}

func compareStrings(item1, item2 string, opts string) (bool, error) {
	// fmt.Printf("%s %s %s (compareStrings)\n", item1, common, item2)
	switch opts {
	case "==", "=":
		return item1 == item2, nil
	case "!=":
		return item1 != item2, nil
	case ">":
		return strings.Compare(item1, item2) > 0, nil
	case ">=":
		return strings.Compare(item1, item2) >= 0, nil
	case "<":
		return strings.Compare(item1, item2) < 0, nil
	case "<=":
		return strings.Compare(item1, item2) <= 0, nil
	}
	return false, fmt.Errorf("%s: unknown operation", opts)
}

func compareNumbers(item1, item2 string, opts string) (bool, error) {
	// fmt.Printf("%s %s %s (compareNumbers)\n", item1, common, item2)
	n1, err1 := strconv.Atoi(item1)
	if err1 != nil {
		return false, err1
	}
	n2, err2 := strconv.Atoi(item2)
	if err2 != nil {
		return false, err2
	}
	switch opts {
	case "-eq":
		return n1 == n2, nil
	case "-ne":
		return n1 != n2, nil
	case "-lt":
		return n1 < n2, nil
	case "-le":
		return n1 <= n2, nil
	case "-gt":
		return n1 > n2, nil
	case "-ge":
		return n1 >= n2, nil
	}
	return false, fmt.Errorf("%s: unknown operation", opts)
}

func compareFiles(file1, file2 string, opts string) (bool, error) {
	// fmt.Printf("%s %s %s (compareFiles)\n", file1, common, file2)
	stat1, err1 := os.Stat(file1)
	stat2, err2 := os.Stat(file2)
	// err1 == nil: file1 is exists
	switch opts {
	case "-nt":
		return err1 == nil && stat1.ModTime().After(stat2.ModTime()), nil
	case "-ot":
		return err1 == nil && stat1.ModTime().Before(stat2.ModTime()), nil
	case "-ef":
		return err1 == nil && err2 == nil && os.SameFile(stat1, stat2), nil
	}
	return false, fmt.Errorf("%s: unknown operation", opts)
}

func isIn(str string, set []string) bool {
	for _, item := range set {
		if str == item {
			return true
		}
	}
	return false
}

func (e *Expression) compare(item1, item2 string, opts string) (bool, error) {
	// fmt.Printf("compare(%s, %s, %s)\n", item1, item2, common)
	str1 := replaceEnv(e, item1)
	str2 := replaceEnv(e, item2)
	if isIn(opts, stringOpts) {
		return compareStrings(str1, str2, opts)
	} else if isIn(opts, numberOpts) {
		return compareNumbers(str1, str2, opts)
	} else if isIn(opts, fileOpts) {
		return compareFiles(str1, str2, opts)
	}
	return true, nil
}

func replaceEnv(e *Expression, str string) string {
	result := str
	for _, env := range e.envs {
		result = strings.ReplaceAll(result, fmt.Sprintf("$%s", env.key), env.value)
		//if result != str {
		//	fmt.Printf("  \"%s\".replace($%s, %s) -> %s\n", str, env.key, env.value, result)
		//}
	}
	//fmt.Printf("replaceEnv(%s) -> %s\n", original, result)
	return result
}

func (e *Expression) Term() (bool, error) {
	c, _ := e.current()
	if strings.HasPrefix(c, "-") {
		target, ok := e.next()
		if !ok {
			fmt.Errorf("%s: parse error", e.Original)
		}
		return operation(c, target)
	}
	opts, ok := e.next()
	if !ok {
		return c != "", nil
	}
	item2, ok := e.next()
	return e.compare(c, item2, opts)
}
