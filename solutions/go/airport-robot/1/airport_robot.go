package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	language := greeter.LanguageName()
	greeting := greeter.Greet(visitorName)
	return fmt.Sprintf("I can speak %s: %s", language, greeting)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}