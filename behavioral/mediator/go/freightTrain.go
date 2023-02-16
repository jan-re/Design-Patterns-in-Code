package main

import "fmt"

type freightTrain struct {
	mediator
}

func (ft *freightTrain) arrive() {
	if ft.mediator.canArrive(ft) {
		fmt.Println("FreightTrain: Arrived")
	} else {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
	}
}

func (ft *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	ft.mediator.notifyAboutDeparture()
}

func (ft *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	ft.arrive()
}
