package main

func main() {
	stationMediator := newStationMediator()

	passengerTrain := &passengerTrain{
		mediator: stationMediator,
	}
	freightTrain := &freightTrain{
		mediator: stationMediator,
	}

	// Communication between trains is handled by the mediator. They never communicate directly.
	passengerTrain.arrive()

	// Freight train cannot arrive, as the station is occupied. It is added into mediator's queue.
	freightTrain.arrive()

	// Passenger train departs. It is removed from the queue.
	// Mediator communicates with the freight train and permits its arrival.
	passengerTrain.depart()
}
