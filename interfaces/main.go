package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

type info interface {
	read(int) string
}

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   info
}

func main() {

	engBot := englishBot{}
	spaBot := spanishBot{}

	engBot.printGreeting()
	spaBot.printGreeting()
	printGreeting(engBot)
	printGreeting(spaBot)
	fmt.Println("")

	testPerson := person{
		firstName: "Hüseyin",
		lastName:  "Büyükdere",
		contact: contactInfo{
			email:   "emailt",
			zipCode: 555,
		},
	}

	fmt.Println(testPerson)
}

func (englishBot) getGreeting() string {
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot englishBot) printGreeting() {
	fmt.Println(engBot.getGreeting())
}

func (spaBot spanishBot) printGreeting() {
	fmt.Println(spaBot.getGreeting())
}

func (contactInfo) read(a int) string {

	return "aaaa"
}
