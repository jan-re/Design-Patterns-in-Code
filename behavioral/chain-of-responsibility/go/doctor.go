package main

import "fmt"

type doctor struct {
	nextHandler department
}

func (d *doctor) process(p *patient) {
	if p.doctorDone {
		fmt.Println("Doctor already examined the patient..")
		d.nextHandler.process(p)
	} else {
		fmt.Println("Doctor healing the patient.")
		p.doctorDone = true
		d.nextHandler.process(p)
	}
}

func (d *doctor) setNext(next department) {
	d.nextHandler = next
}
