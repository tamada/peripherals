package peripherals

import "testing"

func TestVersion(t *testing.T) {
	if VERSION != "0.9.1" {
		t.Errorf("mismatch version wont %s, but got %s", "0.9.1", VERSION)
	}
	wontVersion := "tamada/peripherals version 0.9.1"
	gotVersion := Version("tamada/peripherals")
	if wontVersion != gotVersion {
		t.Errorf("mismatch version information wont %s, but got %s", wontVersion, gotVersion)
	}
}
