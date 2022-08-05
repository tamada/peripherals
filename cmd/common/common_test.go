package common

import "testing"

func TestValidate(t *testing.T) {
	opts1 := Options{Lines: 0, Bytes: 0, Keyword: "", Predicate: "", NoHeader: false}
	if err := opts1.Validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
	opts2 := Options{Lines: 10, Bytes: 10, Keyword: "", Predicate: "", NoHeader: false}
	if err := opts2.Validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
	opts3 := New()
	if err := opts3.Validate(); err == nil {
		t.Errorf("specifying no option causes an error")
	}
}
