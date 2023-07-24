package main

type game struct {
	terrorists        []*Player
	CounterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		CounterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerroris(dressType string) {
	player := newPlayer("CT", dressType)
	c.CounterTerrorists = append(c.CounterTerrorists, player)
	return
}
