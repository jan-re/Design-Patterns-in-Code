package main

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
