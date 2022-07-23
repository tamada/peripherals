package shell

import (
	"testing"
)

func TestNewWithEmptyString(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Errorf("the empty string for New should be error")
	}
}

//func TestEvalStaticPredicate(t *testing.T) {
//	evaluator, err := New("a5 = a5")
//	if err != nil {
//		return
//	}
//	if !evaluator.Eval("a5") {
//		t.Errorf("the result of evaluator did not match, want true, got false")
//	}
//}
//
//func TestEval(t *testing.T) {
//	evaluator, err := New("PTAKER_LINE = a4")
//	if err != nil {
//		return
//	}
//	if !evaluator.Eval("a4") {
//		t.Errorf("the result of evaluator did not match, want true, got false")
//	}
//}

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
