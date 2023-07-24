package main

type NoarmalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NoarmalBuilder {
	return &NoarmalBuilder{}
}

func (b *NoarmalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NoarmalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NoarmalBuilder) setNumFloor() {
	b.floor = 2
}
func (b *NoarmalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
