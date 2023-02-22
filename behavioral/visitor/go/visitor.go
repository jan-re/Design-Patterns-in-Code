package main

type visitor interface {
	visitCircle(*circle)
	visitSquare(*square)
}
