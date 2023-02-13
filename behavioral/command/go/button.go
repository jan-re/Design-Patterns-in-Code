package main

// This button type is an invoker.
// It stores the potential commands and can trigger their executions.
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
