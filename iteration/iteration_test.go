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

func TestCaesarShift(t *testing.T) {
	out := CaesarShift('a', 2)
	expected := 'c'
	if out != expected {
		t.Errorf("Expected %q got %q", expected, out)
	}
}

func TestCaesarWord(t *testing.T) {
	out := CaesarWord("cat", 2)
	expected := "ecv"
	if out != expected {
		t.Errorf("Expected %q got %q", expected, out)
	}
}

func TestRmWhitespace(t *testing.T) {
	out := RmWhitespace("Salami snack plate")
	expected := "salamisnackplate"
	if expected != out {
		t.Errorf("Expected %q got %q", expected, out)
	}
}

func TestVigenere(t *testing.T) {
	out := Vigenere("Salami snack plate", "ratmilk")
	expected := "jaemutceavwxwkke"
	if expected != out {
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
