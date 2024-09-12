package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}
  prefix := greetingPrefix(language)
	return prefix + name + "!"
}

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	}
  return prefix  
}

func main() {
	fmt.Println(Hello("kkdashu", ""))
}
