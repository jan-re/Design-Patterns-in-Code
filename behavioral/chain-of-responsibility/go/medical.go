package main

import "fmt"

type medical struct {
	nextHandler department
}

func (m *medical) process(p *patient) {
	if p.medicalDone {
		fmt.Println("Initial medical check already done.")
		m.nextHandler.process(p)
	} else {
		fmt.Println("Medical examining the patient.")
		p.medicalDone = true
		m.nextHandler.process(p)
	}
}

func (m *medical) setNext(next department) {
	m.nextHandler = next
}
