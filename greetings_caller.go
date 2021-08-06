package main

import (
	"fmt"
	"log"

	GreetingsModule "github.com/Dmaddu/greetings_module"
)

func main() {
	message, err := GreetingsModule.Hello("Durga")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	message, err = GreetingsModule.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
