package main

import (
	"femBasics/imports"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	newTicket := imports.Ticket{
		ID:    1,
		Event: "New Year's Eve Party",
	}

	newTicket.PrintEvent()
	fmt.Println(newTicket)
}
