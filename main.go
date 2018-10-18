package main

import (
	"box_upload/config"
	"crypto/rsa"
	"log"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	privKeyPath string
	verifyKey   *rsa.PublicKey
	signKey     *rsa.PrivateKey
)

func main() {
	err := config.ReadConfiguration()
	fatal(err)
	privKeyPath = config.PrivateKey
	createdToken, err := Assertion()
	fatal(err)
	getToken(createdToken, config.ClientID, config.ClientSecret, true)
	return
}
