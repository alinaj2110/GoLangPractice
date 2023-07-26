package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishbot struct {
}

type spanishbot struct {
}

func main() {
	eb := englishbot{}
	sb := spanishbot{}

	printbot(eb)
	printbot(sb)
}

func (englishbot) getGreeting() string {
	return "Hello"
}

func (spanishbot) getGreeting() string {
	return "Hola"
}

func printbot(b bot) {
	fmt.Println(b.getGreeting())
}
