package main

import "fmt"

type bot interface {
	getGreeting() string
}

type spanishBot struct{}
type englishBot struct{}

func main() {

	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

// This does not work...
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// This does work!
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	return ("Hola!")
}

func (eb englishBot) getGreeting() string {
	return ("Hi there!")
}
