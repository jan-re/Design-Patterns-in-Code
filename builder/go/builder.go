package main

import "errors"

type IBuilder interface {
	AddDoor(*ConstructionSite)
	AddFloor(*ConstructionSite)
	AddWindow(*ConstructionSite)
}

func getBuilder(t string) (IBuilder, error) {
	if t == "house" {
		b := HouseBuilder{}
		return b, nil
	} else if t == "castle" {
		b := CastleBuilder{}
		return b, nil
	} else {
		return nil, errors.New("Invalid builder choice")
	}
}

type HouseBuilder struct {
}

func (b HouseBuilder) AddDoor(cs *ConstructionSite) {
	cs.Door = "Door for a standard house"
}

func (b HouseBuilder) AddFloor(cs *ConstructionSite) {
	cs.Floor = "Your average house floor"
}

func (b HouseBuilder) AddWindow(cs *ConstructionSite) {
	cs.Window = "Standard house window"
}

type CastleBuilder struct {
}

func (b CastleBuilder) AddDoor(cs *ConstructionSite) {
	cs.Door = "Sturdy castle gate"
}

func (b CastleBuilder) AddFloor(cs *ConstructionSite) {
	cs.Floor = "Stone castle floor"
}

func (b CastleBuilder) AddWindow(cs *ConstructionSite) {
	cs.Window = "Arrow slit in the wall"
}

type ConstructionSite struct {
	Door   string
	Floor  string
	Window string
}
