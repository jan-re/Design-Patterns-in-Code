package main

import "fmt"

func main() {
	TerroristSkinType := "tSkin"
	CounterTerroristSkinType := "ctSkin"

	game := newGame()

	game.addCounterTerrorist(CounterTerroristSkinType)
	game.addCounterTerrorist(CounterTerroristSkinType)
	game.addCounterTerrorist(CounterTerroristSkinType)

	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)

	// Retrieve the SkinFactorySingleInstance to prove that only two instances of the Skin types exist
	skinFactoryInstance := getSkinFactorySingleInstance()

	for skinType, skin := range skinFactoryInstance.skinMap {
		fmt.Printf("SkinColorType: %s\nSkinColor: %s\n", skinType, skin.getColor())
	}
}
