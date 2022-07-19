package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tamada/peripherals/eval"
)

type ShellEvaluator struct {
	shellName string
	args      []string
}

func NewWithShell(shellName string) (eval.Evaluator, error) {
	if shellName == "" {
		return nil, fmt.Errorf("empty shellName does not allow")
	}
	args := []string{}
	if isIn(shellName, []string{"bash", "zsh", "sh"}) {
		args = []string{"-c"}
	}
	return &ShellEvaluator{shellName: shellName, args: args}, nil
}

func New() (eval.Evaluator, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return nil, fmt.Errorf("SHELL: environment value was not found")
	}
	return NewWithShell(shell)
}

func isIn(value string, set []string) bool {
	for _, v := range set {
		if strings.Contains(v, value) {
			return true
		}
	}
	return false
}

func (se *ShellEvaluator) Eval(predicate, line string) bool {
	cmd := setupCmd(se, predicate, line)
	err := cmd.Run()
	return err == nil
}

func setupCmd(se *ShellEvaluator, predicate, line string) *exec.Cmd {
	newArgs := []string{}
	newArgs = append(newArgs, se.args...)
	newArgs = append(newArgs, strings.Split(predicate, "")...)
	cmd := exec.Command(se.shellName, newArgs...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("CLINE=%s", line))
	return cmd
}
