package main

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (sc *SecurityCode) checkSecurityCode(code int) error {
	if sc.code != code {
		return errors.New("Provided security code is not correct. Transaction denied.")
	} else {
		fmt.Println("Security code accepted.")
		return nil
	}
}
