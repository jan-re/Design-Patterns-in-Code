package main

// Flyweight interface
type Skin interface {
	getColor() string
}

// Concrete flyweight instance and associated functions
// Struct should only be created once and pointed to from that moment.
type TSkin struct {
	color string
}

func newTSkin() *TSkin {
	return &TSkin{
		color: "beige",
	}
}

func (t *TSkin) getColor() string {
	return t.color
}

// Concrete flyweight instance and associated functions
// Struct should only be created once and pointed to from that moment.
type CTSkin struct {
	color string
}

func newCTSkin() *TSkin {
	return &TSkin{
		color: "black",
	}
}

func (ct *CTSkin) getColor() string {
	return ct.color
}
