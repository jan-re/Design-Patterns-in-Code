package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildFullProduct() {
	d.builder.AddDoor()
	d.builder.AddWall()
	d.builder.AddWindow()
}

func (d *Director) getHouse() House {
	return d.builder.getHouse()
}
