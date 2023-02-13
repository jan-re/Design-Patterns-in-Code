package main

// The shared interface for receivers.
type device interface {
	on()
	off()
}
