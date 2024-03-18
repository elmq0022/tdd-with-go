package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const norwegianHelloPrefix = "Hei, "

const french = "French"
const norwegian = "Norwegian"
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
}
