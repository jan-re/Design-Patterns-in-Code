package main

import "fmt"

// This is the initial step, before decoration.
type Sender struct {
}

func (s *Sender) dispatch() string {
	message := "This is a generic message. Hello.\n"
	fmt.Printf("Sending: %s\n", message)

	return message
}
