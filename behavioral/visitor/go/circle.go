package main

type circle struct {
	shapeType string
}

func newCircle() *circle {
	return &circle{
		shapeType: "Circle",
	}
}

func (c *circle) getType() string {
	return c.shapeType
}

func (c *circle) accept(v visitor) {
	v.visitCircle(c)
}
