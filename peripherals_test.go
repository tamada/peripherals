package peripherals

import "testing"

func TestVersion(t *testing.T) {
	if VERSION != "0.9.0" {
		t.Errorf("mismatch version wont %s, but got %s", "", VERSION)
	}
	wontVersion := "tamada/peripherals version 0.9.0"
	gotVersion := Version("tamada/peripherals")
	if wontVersion != gotVersion {
		t.Errorf("mismatch version information wont %s, but got %s", wontVersion, gotVersion)
	}
}
