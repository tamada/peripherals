package shell

import (
	"fmt"
	"github.com/tamada/peripherals/eval"
	"os"
	"os/exec"
	"strings"
)

type Evaluator struct {
	shellName string
	args      []string
}

func New(predicate string) (eval.Evaluator, error) {
	if predicate == "" {
		return nil, fmt.Errorf("empty predicate does not allow")
	}
	args := strings.Split(predicate, " ")
	newArgs := []string{}
	if len(args) >= 1 {
		newArgs = args[1:]
	}
	return &Evaluator{shellName: args[0], args: newArgs}, nil
}

func (se *Evaluator) Eval(line string) bool {
	cmd := setupCmd(se, line)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return cmd.ProcessState.ExitCode() == 0
}

func setupCmd(se *Evaluator, line string) *exec.Cmd {
	cmd := exec.Command(se.shellName, se.args...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("CLINE=%s", line))
	return cmd
}
