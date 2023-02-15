package main

type stationMediator struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationMediator() *stationMediator {
	return &stationMediator{
		isPlatformFree: true,
	}
}

func (sm *stationMediator) canArrive(t Train) bool {
	// Continue from here
}
