package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	BoxAppSettings struct {
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
		AppAuth      struct {
			PublicKeyID string `json:"publicKeyID"`
			PrivateKey  string `json:"privateKey"`
			Passphrase  string `json:"passphrase"`
		} `json:"appAuth"`
	} `json:"boxAppSettings"`
	EnterpriseID string `json:"enterpriseID"`
}

var (
	ClientID     string
	ClientSecret string
	AppAuth      string
	PublicKeyID  string
	PrivateKey   string
	Passphrase   string
	EnterpriseID string
)
var cfg Config

func ReadConfiguration() error {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into which we defined above
	json.Unmarshal(byteValue, &cfg)

	// we set our variables here

	ClientID = cfg.BoxAppSettings.ClientID
	ClientSecret = cfg.BoxAppSettings.ClientSecret
	Passphrase = cfg.BoxAppSettings.AppAuth.Passphrase

	PublicKeyID = cfg.BoxAppSettings.AppAuth.PublicKeyID
	PrivateKey = cfg.BoxAppSettings.AppAuth.PrivateKey
	EnterpriseID = cfg.EnterpriseID

	return nil
}
