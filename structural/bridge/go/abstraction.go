package main

type Abstraction struct {
	Implementation
}

func (a *Abstraction) runOperation() string {
	return a.Implementation.implementOperation()
}
