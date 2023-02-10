package main

import "fmt"

type WalletFacade struct {
	account *Account
	code    *SecurityCode
	wallet  *Wallet
	log     *Log
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	wf := WalletFacade{
		account: newAccount(accountID),
		code:    newSecurityCode(code),
		wallet:  newWallet(),
		log:     &Log{},
	}

	return &wf
}

func (wf *WalletFacade) securityCheck(accountID string, securityCode int) error {
	err := wf.account.checkAccountId(accountID)
	if err != nil {
		return err
	}

	err = wf.code.checkSecurityCode(securityCode)
	if err != nil {
		return err
	}

	return nil
}

func (wf *WalletFacade) depositFunds(accountID string, securityCode int, amount int) error {
	fmt.Println("Checking credentials...")
	err := wf.securityCheck(accountID, securityCode)
	if err != nil {
		return err
	}

	wf.wallet.addBalance(amount)
	wf.log.logTransaction(accountID, "deposit", amount)

	return nil
}

func (wf *WalletFacade) withdrawFunds(accountID string, securityCode int, amount int) error {
	fmt.Println("Checking credentials...")
	err := wf.securityCheck(accountID, securityCode)
	if err != nil {
		return err
	}

	err = wf.wallet.withdrawBalance(amount)
	if err != nil {
		return err
	}

	wf.log.logTransaction(accountID, "withdrawal", amount)

	return nil
}
