package main

func main() {
	stationManager := newStationManager()

	PassengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	PassengerTrain.arrive()
	freightTrain.arrive()
	PassengerTrain.depart()
}
