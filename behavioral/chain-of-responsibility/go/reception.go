package main

import "fmt"

type reception struct {
	nextHandler department
}

func (r *reception) process(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.nextHandler.process(p)
	} else {
		fmt.Println("Reception registering patient")
		p.registrationDone = true
		r.nextHandler.process(p)
	}
}

func (r *reception) setNext(next department) {
	r.nextHandler = next
}
