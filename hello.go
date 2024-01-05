package main

import "fmt"

func Hello(name string) string {
	return "Yoo, " + name
}

func main() {
	fmt.Println(Hello("x"))
}
