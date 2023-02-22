package main

func main() {
	square := newSquare()
	circle := newCircle()

	shapes := []shape{square, circle}

	xml := xmlVisitor{}

	for _, shape := range shapes {
		shape.accept(xml)
	}
}
