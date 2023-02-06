package main

import "fmt"

type Sender struct {
}

func (s *Sender) dispatch() string {
	message := "This is a generic message. Hello.\n"
	fmt.Printf("Sending: %s\n", message)

	return message
}
