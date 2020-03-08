package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	eb.getGreetings()
	sb.getGreetings()

	printGreetings(eb)
	printGreetings(sb)
}

func (eb englishBot) getGreetings() string {
	return "Hi, there!"
}

func (sb spanishBot) getGreetings() string {
	return "Hola!"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
