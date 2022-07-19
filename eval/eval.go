package eval

type Evaluator interface {
	Eval(condition string, line string) bool
}
