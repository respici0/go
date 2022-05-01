package main

import "fmt"

type bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func main() {

	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (EnglishBot) getGreeting() string {
	return "Hello!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}
