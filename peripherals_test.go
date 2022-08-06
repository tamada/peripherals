package peripherals

import "testing"

func TestVersion(t *testing.T) {
	if VERSION != "" {
		t.Errorf("mismatch version wont %s, but got %s", "", VERSION)
	}
	wontVersion := "tamada/peripherals "
	gotVersion := Version()
	if wontVersion != gotVersion {
		t.Errorf("mismatch version information wont %s, but got %s", wontVersion, gotVersion)
	}
}
