package pskip

import (
	"github.com/tamada/peripherals/ptake"
)

type Skipper struct {
	taker ptake.Taker
}

func New(kind ptake.TakerType, data ptake.InputData) (ptake.Taker, error) {
	taker, err := ptake.New(kind, data)
	return &Skipper{taker: taker}, err
}

func (skipper *Skipper) IsLine() bool {
	return skipper.taker.IsLine()
}

func (skipper *Skipper) Reset() {
	skipper.taker.Reset()
}

func (skipper *Skipper) TakeLine(data string) bool {
	return !skipper.taker.TakeLine(data)
}

func (skipper *Skipper) TakeByte(data byte) bool {
	return !skipper.taker.TakeByte(data)
}
