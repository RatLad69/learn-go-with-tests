package main

import "fmt"

const englishHelloPrefix = "Yoo, "

func Hello(name string) string {
	if name == "" {
		name = "bungholio"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("x"))
}
