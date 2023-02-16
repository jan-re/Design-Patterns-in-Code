package main

import "fmt"

type passengerTrain struct {
	mediator
}

func (ft *passengerTrain) arrive() {
	if ft.mediator.canArrive(ft) {
		fmt.Println("PassengerTrain: Arrived")
	} else {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
	}
}

func (ft *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	ft.mediator.notifyAboutDeparture()
}

func (ft *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted")
	ft.arrive()
}
