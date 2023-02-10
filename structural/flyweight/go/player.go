package main

import "log"

type Player struct {
	skin       Skin
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, skinType string) *Player {
	skin, err := getSkinFactorySingleInstance().getSkinByType(skinType)
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		playerType: playerType,
		skin:       skin,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
