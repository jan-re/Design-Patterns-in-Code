package main

// Concrete command fulfilling the command interface.
type switchOn struct {
	device device
}

func (so *switchOn) execute() {
	so.device.on()
}
