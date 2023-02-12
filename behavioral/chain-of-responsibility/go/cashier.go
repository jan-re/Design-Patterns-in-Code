package main

import "fmt"

// Intended to be the last handler in the chain.
type cashier struct {
	nextHandler department
}

func (c *cashier) process(p *patient) {
	if p.cashierDone {
		fmt.Println("Payment done.")
	} else {
		fmt.Println("Billing the patient.")
		p.cashierDone = true
	}
}

func (c *cashier) setNext(next department) {
	c.nextHandler = next
}
