package iteration

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Repeater(ch string, x int) string {
	var repeated string
	for i := 0; i < x; i++ {
		repeated += ch
	}
	return repeated
}

func LeftPad(inString string, size int) string {
	// JUSTICE FOR LEFTPAD
	if len(inString) >= size {
		return inString
	}
	preString := Repeater(" ", size-len(inString))
	return preString + inString
}

func CaesarShift(input rune, shift int) rune {
	shifted := int(input) + shift
	if shifted > 'z' {
		return rune(shifted - 26)
	} else if shifted < 'a' {
		return rune(shifted + 26)
	}
	return rune(shifted)
}

func CaesarWord(input string, shift int) string {
	shiftedWord := ""
	for _, ch := range input {
		shiftedWord += string(CaesarShift(rune(ch), shift))
	}
	return shiftedWord
}

func RmWhitespace(input string) string {
	processedString := strings.ToLower(input)
	var b strings.Builder
	b.Grow(len(processedString))
	for _, ch := range processedString {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func Vigenere(input string, key string) string {
	processedString := RmWhitespace(input)
	vString := ""
	keyLength := int(utf8.RuneCountInString(key))
	for i, ch := range processedString {
		vString += string(CaesarShift(rune(ch), int(key[i%keyLength])-97))
	}
	return vString
}
