package main

import "fmt"

// Decorators contain the original structure as a field.
type Logger struct {
	Communique
}

func (l *Logger) dispatch() string {
	message := l.Communique.dispatch()
	fmt.Printf("Message logged: %s\n", message)

	return message
}
