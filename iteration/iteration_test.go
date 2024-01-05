package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	out := Repeater("a", 5)
	expected := "aaaaa"
	if out != expected {
		t.Errorf("Expected %q got %q", expected, out)
	}
}

func TestLeftPad(t *testing.T) {
	out := LeftPad("Q", 4)
	expected := "   Q"
	if out != expected {
		t.Errorf("Expected %q got %q", expected, out)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}

func BenchmarkLeftPad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeftPad("QQ", 7)
	}
}

func ExampleRepeater() {
	out := Repeater("q", 3)
	fmt.Println(out)
	// Output: qqq
}

func ExampleLeftPad() {
	out := LeftPad("a", 3)
	fmt.Println(out)
	// Output:   a
}
