package main

import (
	"box_upload/config"
	dorand "box_upload/rand"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Assertion() (string, error) {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword(signBytes, config.Passphrase)
	fatal(err)

	token := jwt.New(jwt.SigningMethodRS256)
	// set kid
	token.Header["kid"] = config.PublicKeyID
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["iss"] = config.ClientID
	claims["sub"] = config.EnterpriseID
	claims["aud"] = "https://api.box.com/oauth2/token"
	claims["box_sub_type"] = "enterprise"
	claims["jti"] = dorand.StringWithCharset(16)
	claims["exp"] = time.Now().Add(time.Second * 60).Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	fatal(err)

	return tokenString, err
}
