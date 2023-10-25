package main

import (
	"fmt"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return HelloText(language) + name
}

func HelloText(language string) (prefix string) {
	switch language {
	case "japanese":
		prefix = "こんにちは,"
	case "english":
		prefix = "Hello,"
	case "spanish":
		prefix = "Hola,"
	default:
		prefix = "Hello,"
	}
	return
}

func main() {
	fmt.Println(Hello("ryo", "japanese"))
}
