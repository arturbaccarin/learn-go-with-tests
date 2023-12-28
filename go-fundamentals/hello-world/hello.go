package helloworld

import "fmt"

const (
	spanish               = "Spanish"
	portuguese            = "Portuguese"
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	portugueseHelloPrefix = "Olá, "
)

func Hello() string {
	return "Hello, world"
}

func HelloFriend(name, language string) string {
	var prefix string

	if name == "" {
		name = "World"
	}

	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello())
}
