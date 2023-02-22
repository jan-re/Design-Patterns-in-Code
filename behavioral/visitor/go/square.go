package main

type square struct {
	shapeType string
}

func newSquare() *square {
	return &square{
		shapeType: "Square",
	}
}

func (s *square) getType() string {
	return s.shapeType
}

func (s *square) accept(v visitor) {
	v.visitSquare(s)
}
