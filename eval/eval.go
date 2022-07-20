package eval

type Evaluator interface {
	Eval(line string) bool
}
