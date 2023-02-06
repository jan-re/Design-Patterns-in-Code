package main

import "fmt"

type Logger struct {
	Communique
}

func (l *Logger) dispatch() string {
	message := l.Communique.dispatch()
	fmt.Printf("Message logged: %s\n", message)

	return message
}
