package rectangle_test

import (
	"testing"

	"github.com/rossnoble/gotour/04-methods/rectangle"
)

func TestRectangle(t *testing.T) {
	r := Rectangle{3, 5}

	if r.Width() != 3 {
		t.Errorf("expected 3")
	}

	if r.Height() != 5 {
		t.Errorf("expected 5")
	}
}
