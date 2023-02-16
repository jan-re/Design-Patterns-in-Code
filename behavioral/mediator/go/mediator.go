package main

// Generic mediator interface
type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}
