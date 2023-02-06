package main

import "fmt"

func main() {
	message := Sender{}

	logger := Logger{
		&message,
	}

	encrypter := Encrypter{
		&logger,
	}

	fmt.Printf("Encrypted message: %s\n", encrypter.dispatch())
}
