package main

import "errors"

type IBuilder interface {
	AddDoor()
	AddWall()
	AddWindow()
	getHouse() House
}

func getBuilder(t string) (IBuilder, error) {
	if t == "house" {
		return newHouseBuilder(), nil
	} else if t == "castle" {
		return newCastleBuilder(), nil
	} else {
		return nil, errors.New("Invalid builder choice")
	}
}
