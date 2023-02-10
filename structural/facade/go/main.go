package main

import (
	"fmt"
	"log"
)

func main() {
	accountID := "account123"
	securityCode := 1234

	facade := newWalletFacade(accountID, securityCode)

	facade.depositFunds(accountID, securityCode, 500)
	err := facade.withdrawFunds(accountID, securityCode, 300)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Attempting to overdraw...")
	err = facade.withdrawFunds(accountID, securityCode, 1000)
	if err != nil {
		log.Fatal(err)
	}
}

/*
The point of this little application is to prove that while the components
here are many and their relationships varied, the facade allows the client
to interact with them in a very simple, abstracted way.
*/
