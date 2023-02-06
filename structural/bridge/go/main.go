package main

import "fmt"

func main() {
	impA := ImplementationA{}
	abs := Abstraction{
		Implementation: &impA,
	}

	fmt.Println(abs.runOperation())

	impB := ImplementationB{}
	abs = Abstraction{
		Implementation: &impB,
	}

	fmt.Println(abs.runOperation())
}

// Alternative example: https://refactoring.guru/design-patterns/bridge/go/example#example-0
