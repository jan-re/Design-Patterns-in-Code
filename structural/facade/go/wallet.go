package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) addBalance(amount int) {
	w.balance = w.balance + amount

	fmt.Println("Funds deposited.")
}

func (w *Wallet) withdrawBalance(amount int) error {
	if amount > w.balance {
		return errors.New(fmt.Sprintf("Your current balance is %d. Withdrawing %d would push your account into negative numbers. Transaction denied.", w.balance, amount))
	} else {
		w.balance = w.balance - amount

		fmt.Println("Funds withdrawn.")
		return nil
	}
}
