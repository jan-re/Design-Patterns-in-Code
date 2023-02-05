package main

type CastleBuilder struct {
	doorType   string
	wallType   string
	windowType string
}

func newCastleBuilder() *CastleBuilder {
	return &CastleBuilder{}
}

/*
First, and most important, does the method need to modify the receiver? If it does, the receiver must be a pointer. (Slices and maps act as references, so their story is a little more subtle, but for instance to change the length of a slice in a method the receiver must still be a pointer.)

In the examples above, if pointerMethod modifies the fields of s, the caller will see those changes, but valueMethod is called with a copy of the caller's argument (that's the definition of passing a value), so changes it makes will be invisible to the caller.
*/

func (b *CastleBuilder) AddDoor() {
	b.doorType = "Mighty castle gate"
}

func (b *CastleBuilder) AddWall() {
	b.wallType = "Stone curtain wall"
}

func (b *CastleBuilder) AddWindow() {
	b.windowType = "Arrow slits"
}

func (b CastleBuilder) getHouse() House {
	return House{
		Door:   b.doorType,
		Wall:   b.wallType,
		Window: b.windowType,
	}
}

type HouseBuilder struct {
	doorType   string
	wallType   string
	windowType string
}

func newHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (b *HouseBuilder) AddDoor() {
	b.doorType = "Door for a standard house"
}

func (b *HouseBuilder) AddWall() {
	b.wallType = "Your average house wall"
}

func (b *HouseBuilder) AddWindow() {
	b.windowType = "Standard house window"
}

func (b HouseBuilder) getHouse() House {
	return House{
		Door:   b.doorType,
		Wall:   b.wallType,
		Window: b.windowType,
	}
}
