package main

import "encoding/base64"

type Encrypter struct {
	Communique
}

// No actual encryption, just a simplification
func (e *Encrypter) dispatch() string {
	plaintext := []byte(e.Communique.dispatch())

	return base64.URLEncoding.EncodeToString(plaintext)
}
