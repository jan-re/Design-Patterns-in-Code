package main

import "fmt"

// Concrete receiver.
type tv struct {
	isRunning bool
}

func (tv *tv) on() {
	if tv.isRunning {
		fmt.Println("TV is already running.")
	} else {
		tv.isRunning = true
		fmt.Println("TV switched on.")
	}
}

func (tv *tv) off() {
	if !tv.isRunning {
		fmt.Println("TV is already off.")
	} else {
		tv.isRunning = false
		fmt.Println("TV switched off.")
	}
}
