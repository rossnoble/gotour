package rectangle_test

import (
	"testing"

	. "."
)

//. "github.com/rossnoble/gotour/04-methods/rectangle"

func TestRectangleWidth(t *testing.T) {
	r := Rectangle{3, 5}

	if r.Width() != 3 {
		t.Errorf("expected 3")
	}
}

func TestRectangleLength(t *testing.T) {
	r := Rectangle{3, 5}
	if r.Length() != 5 {
		t.Errorf("expected 5")
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{3, 5}

	if r.Perimeter() != 16 {
		t.Errorf("expected 16")
	}
}
