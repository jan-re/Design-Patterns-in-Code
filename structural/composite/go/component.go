package main

/*
The expectation is that the Component will contain both files and slices of other components.
*/
type Component interface {
	search(string)
}
