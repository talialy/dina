package tests

import (
	"testing"

	"github.com/talialy/dina/cmd"
)

func TestPkgManagers(t *testing.T) {
	pkg, err := cmd.GetPkg()
	want := "dnf"
	if err != nil {
		t.Fatal(err)
	}

	if pkg != "dnf" {
		t.Errorf("it found %s, when it was really %s", pkg, want)
	}
}
