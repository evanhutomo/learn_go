package mascot_test

import (
	"testing"

	"eh.org/vscode_sample/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Gopher" {
		t.Fatal("wrong mascot")
	}
}
