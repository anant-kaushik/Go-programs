package main

import (
	"fmt"
	"anantkaushik.com/greetings"
)

func main() {
		message := greetings.Hello("Anant")
		fmt.Println(message)
}