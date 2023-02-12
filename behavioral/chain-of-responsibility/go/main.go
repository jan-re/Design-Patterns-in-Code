package main

import "fmt"

func main() {
	patientJohn := patient{
		name:             "John",
		doctorDone:       false,
		registrationDone: false,
		medicalDone:      false,
		cashierDone:      false,
	}

	reception := reception{}

	medical := medical{}
	reception.setNext(&medical)

	doctor := doctor{}
	medical.setNext(&doctor)

	cashier := cashier{}
	doctor.setNext(&cashier)

	reception.process(&patientJohn)

	fmt.Printf("Name: %s\nRegistration done: %v\nMedical done: %v\nDoctor done: %v\nCashier done: %v\n",
		patientJohn.name,
		patientJohn.registrationDone,
		patientJohn.medicalDone,
		patientJohn.doctorDone,
		patientJohn.cashierDone,
	)

}
