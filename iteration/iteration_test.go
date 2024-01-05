package iteration

import "testing"

func TestRepeater(t *testing.T) {
	out := Repeater("a", 5)
	expected := "aaaaa"
	if out != expected {
		t.Errorf("Expected %q got %q", expected, out)
	}
}
