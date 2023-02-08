package main

import (
	"errors"
	"fmt"
)

type Account struct {
	id string
}

func newAccount(id string) *Account {
	return &Account{
		id: id,
	}
}

func (ac *Account) checkAccountId(id string) error {
	if ac.id != id {
		return errors.New("Provided account ID is not correct. Transaction denied.")
	} else {
		fmt.Println("Account verified.")
		return nil
	}
}
