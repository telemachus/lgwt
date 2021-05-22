package main

import "fmt"

func greeting(lang string) string {
	// default to English
	s := "Hello"
	switch lang {
	case "sp":
		s = "¡Hola"
	case "fr":
		s = "Bonjour"
	}

	return s

}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", greeting(lang), name)
}

func main() {
	fmt.Println(Hello("", ""))
}
