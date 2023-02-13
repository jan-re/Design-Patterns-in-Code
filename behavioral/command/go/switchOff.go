package main

// Concrete command fulfilling the command interface.
type switchOff struct {
	device device
}

func (so *switchOff) execute() {
	so.device.off()
}
