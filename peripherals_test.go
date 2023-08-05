package peripherals

import "testing"

func TestVersion(t *testing.T) {
	if VERSION != "1.0.0" {
		t.Errorf("mismatch version wont %s, but got %s", "1.0.0", VERSION)
	}
	wontVersion := "tamada/peripherals version 1.0.0"
	gotVersion := Version("tamada/peripherals")
	if wontVersion != gotVersion {
		t.Errorf("mismatch version information wont %s, but got %s", wontVersion, gotVersion)
	}
}
