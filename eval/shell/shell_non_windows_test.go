//go:build !windows

package shell

import (
	"testing"
)

func TestScriptEval(t *testing.T) {
	evaluator, err := New("../../testdata/eval_script.sh")
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
		return
	}
	if !evaluator.Eval("a1") {
		t.Errorf("the result of evaluator did not match, want true, got false")
	}
}
