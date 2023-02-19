package main

import "fmt"

func main() {
	nvidia := newEshopItem("1080ti")

	customers := []*customer{{id: "Michael"}, {id: "Jan"}, {id: "Jaro"}}

	for _, customer := range customers {
		nvidia.register(customer)
	}

	fmt.Println("------------------------------------------------")

	nvidia.deregister(&customer{id: "Jaro"})

	fmt.Println("------------------------------------------------")

	nvidia.makeAvailable()
}
