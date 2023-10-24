package main

import "fmt"

const helloText = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloText + name
}

func main() {
	fmt.Println(Hello("world"))
}
