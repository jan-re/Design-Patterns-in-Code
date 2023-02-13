package main

// The shared interface for all commands.
type command interface {
	execute()
}
