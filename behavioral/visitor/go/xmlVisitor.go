package main

import "fmt"

type xmlVisitor struct{}

func (v xmlVisitor) visitCircle(c *circle) {
	fmt.Println("Converting shape Circle into XML...")
}

func (v xmlVisitor) visitSquare(s *square) {
	fmt.Println("Converting shape Square into XML...")
}
