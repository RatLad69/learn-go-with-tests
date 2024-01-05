package main

import (
	"fmt"
)

const englishHelloPrefix = "Yoo, "
const klingonHelloPrefix = "nuqneH, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "bungholio"
	}
	prefix := englishHelloPrefix
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Klingon":
		prefix = klingonHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("x", "English"))
}
