package main

import "fmt"

const (
	french = "French"
	spanish = "Spanish"
	polish = "Polish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	polishHelloPrefix = "No elo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return 
}

func main() {
	fmt.Println(Hello("Mariusz", ""))
}
