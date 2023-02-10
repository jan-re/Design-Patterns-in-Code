package main

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        []*Player{},
		counterTerrorists: []*Player{},
	}
}

func (c *Game) addTerrorist(skinType string) {
	player := newPlayer("T", skinType)
	c.terrorists = append(c.terrorists, player)
}

func (c *Game) addCounterTerrorist(skinType string) {
	player := newPlayer("CT", skinType)
	c.counterTerrorists = append(c.counterTerrorists, player)
}
