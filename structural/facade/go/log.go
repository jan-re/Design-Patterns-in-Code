package main

import "fmt"

type Log struct {
	loggedTransactions []Transaction
}

func newLog() *Log {
	return &Log{
		loggedTransactions: []Transaction{},
	}
}

func (l *Log) logTransaction(accountID string, transactionType string, amount int) {
	l.loggedTransactions = append(l.loggedTransactions, Transaction{
		accountID:       accountID,
		transactionType: transactionType,
		amount:          amount,
	})

	fmt.Println("This transaction has been logged.")
}
