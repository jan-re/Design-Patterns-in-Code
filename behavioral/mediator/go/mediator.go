package main

// Generic mediator interface
type mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
