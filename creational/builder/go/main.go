package main

import (
	"fmt"
	"log"
)

func main() {
	houseBuilder, err := getBuilder("house")
	if err != nil {
		log.Fatal(err)
	}
	castleBuilder, err := getBuilder("castle")
	if err != nil {
		log.Fatal(err)
	}

	director := newDirector(houseBuilder)

	director.buildFullProduct()
	house := director.getHouse()

	fmt.Println(house)

	director.setBuilder(castleBuilder)
	director.buildFullProduct()
	castle := director.getHouse()

	fmt.Println(castle)
}
