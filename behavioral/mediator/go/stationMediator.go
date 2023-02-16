package main

type stationMediator struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationMediator() *stationMediator {
	return &stationMediator{
		isPlatformFree: true,
	}
}

func (sm *stationMediator) canArrive(t train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}

	sm.trainQueue = append(sm.trainQueue, t)
	return false
}

func (sm *stationMediator) notifyAboutDeparture() {
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}

	// Separate the first train in queue and update the queue to move by one space.
	firstTrain := sm.trainQueue[0]
	sm.trainQueue = sm.trainQueue[1:]

	firstTrain.permitArrival()
}
