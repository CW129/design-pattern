package main

import "fmt"

func main() {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerroris(CounterTerroristDressType)
	game.addCounterTerroris(CounterTerroristDressType)
	game.addCounterTerroris(CounterTerroristDressType)

	dressFactorySingleInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactorySingleInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

}
