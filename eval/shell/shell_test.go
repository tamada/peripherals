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
