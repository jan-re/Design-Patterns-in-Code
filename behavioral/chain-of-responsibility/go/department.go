package main

type department interface {
	process(*patient)
	setNext(department)
}
