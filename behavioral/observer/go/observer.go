package main

import "fmt"

// The generic update method is needed for the subject to notify its observers.
type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending update to customer ID %s. Item %s is now available.\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
