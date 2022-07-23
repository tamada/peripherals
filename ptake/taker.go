package ptake

import (
	"fmt"
	"github.com/tamada/peripherals/eval"
	"github.com/tamada/peripherals/eval/shell"
	"github.com/tamada/peripherals/ptest"
	"strconv"
	"strings"
)

type TakerType int

const (
	LINE TakerType = iota + 1
	BYTES
	WHILE
	UNTIL
)

type Taker interface {
	TakeLine(data string) bool
	TakeByte(data byte) bool
	IsLine() bool
	Reset()
}

type InputData interface {
	Int() int
	String() string
}

type takerLine struct {
	lines   int
	current int
	finish  bool
}

type takerBytes struct {
	bytes   int
	current int
	finish  bool
}

type takerWhile struct {
	expression []string
	lineCount  int
	finish     bool
}

type takerExecWhile struct {
	evaluator eval.Evaluator
	finish    bool
}

type takerUntil struct {
	keyword string
	finish  bool
}

func NewTaker(t TakerType, data InputData) (Taker, error) {
	switch t {
	case WHILE:
		str := data.String()
		if strings.HasSuffix(str, ".sh") {
			evaluator, err := shell.New(str)
			return &takerExecWhile{evaluator: evaluator, finish: false}, err
		} else {
			return &takerWhile{expression: strings.Split(str, " "), lineCount: 0, finish: false}, nil
		}
	case UNTIL:
		return &takerUntil{keyword: data.String(), finish: false}, nil
	case LINE:
		return &takerLine{lines: data.Int(), current: 0, finish: false}, nil
	case BYTES:
		return &takerBytes{bytes: data.Int(), current: 0, finish: false}, nil
	default:
		return nil, fmt.Errorf("unknown error")
	}
}

func (taker *takerBytes) Reset() {
	taker.current = 0
	taker.finish = false
}

func (taker *takerBytes) IsLine() bool {
	return false
}

func (taker *takerBytes) TakeByte(_ byte) bool {
	if !taker.finish {
		taker.current++
		taker.finish = taker.bytes < taker.current
	}
	return !taker.finish
}

func (taker *takerBytes) TakeLine(_ string) bool {
	return true
}

func (taker *takerLine) Reset() {
	taker.current = 0
	taker.finish = false
}

func (taker *takerLine) IsLine() bool {
	return true
}

func (taker *takerLine) TakeByte(_ byte) bool {
	return true
}

func (taker *takerLine) TakeLine(_ string) bool {
	if taker.finish {
		return !taker.finish
	}
	taker.current++
	if taker.lines < taker.current {
		taker.finish = true
	}
	return !taker.finish
}

func (taker *takerUntil) Reset() {
	taker.finish = false
}

func (taker *takerUntil) IsLine() bool {
	return true
}

func (taker *takerUntil) TakeByte(_ byte) bool {
	return true
}

func (taker *takerUntil) TakeLine(data string) bool {
	if !taker.finish {
		if data == taker.keyword {
			taker.finish = true
		}
	}
	return !taker.finish
}

func (taker *takerWhile) Reset() {
	taker.lineCount = 0
	taker.finish = false
}

func (taker *takerWhile) IsLine() bool {
	return true
}

func (taker *takerWhile) TakeByte(_ byte) bool {
	return true
}

func (taker *takerWhile) TakeLine(data string) bool {
	if !taker.finish {
		taker.lineCount++
		e := ptest.New(taker.expression)
		e.SetEnv("PTAKE_LINE", data)
		e.SetEnv("PTAKE_LINECOUNT", strconv.Itoa(taker.lineCount))
		flag, err := e.Eval()
		// fmt.Printf("%s (PTAKER_LINE: %s): %v\n", e.Original, data, flag)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		if !flag {
			taker.finish = true
		}
	}
	return !taker.finish
}

func (taker *takerExecWhile) Reset() {
	taker.finish = false
}

func (taker *takerExecWhile) IsLine() bool {
	return true
}

func (taker *takerExecWhile) TakeByte(_ byte) bool {
	return true
}

func (taker *takerExecWhile) TakeLine(data string) bool {
	if !taker.finish {
		if !taker.evaluator.Eval(data) {
			taker.finish = true
		}
	}
	return !taker.finish
}
