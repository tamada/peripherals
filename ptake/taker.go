package ptake

import (
	"fmt"

	"github.com/tamada/peripherals/eval"
	"github.com/tamada/peripherals/eval/shell"
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
	predicate string
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
		evaluator, err := shell.New()
		return &takerWhile{predicate: data.String(), evaluator: evaluator, finish: false}, err
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

func (taker *takerBytes) TakeByte(data byte) bool {
	if !taker.finish {
		taker.current++
		taker.finish = taker.bytes < taker.current
	}
	return !taker.finish
}

func (taker *takerBytes) TakeLine(data string) bool {
	return true
}

func (taker *takerUntil) Reset() {
	taker.finish = false
}

func (taker *takerUntil) IsLine() bool {
	return true
}

func (taker *takerUntil) TakeByte(data byte) bool {
	return true
}

func (taker *takerUntil) TakeLine(data string) bool {
	if taker.finish {
		return !taker.finish
	}
	if data == taker.keyword {
		taker.finish = true
		return true // print the matched line
	}
	return !taker.finish
}

func (taker *takerLine) Reset() {
	taker.current = 0
	taker.finish = false
}

func (tl *takerLine) IsLine() bool {
	return true
}

func (tl *takerLine) TakeByte(data byte) bool {
	return true
}

func (tl *takerLine) TakeLine(data string) bool {
	if tl.finish {
		return !tl.finish
	}
	tl.current++
	if tl.lines < tl.current {
		tl.finish = true
	}
	return !tl.finish
}

func (taker *takerWhile) Reset() {
	taker.finish = false
}

func (taker *takerWhile) IsLine() bool {
	return true
}

func (taker *takerWhile) TakeByte(data byte) bool {
	return true
}

func (taker *takerWhile) TakeLine(data string) bool {
	if taker.finish {
		return !taker.finish
	}
	if taker.evaluator.Eval(taker.predicate, data) {
		taker.finish = true
		return true
	}
	return !taker.finish
}
