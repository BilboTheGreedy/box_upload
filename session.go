package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getToken(token string, client_id string, client_secret string, debug bool) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if debug == true {
		fmt.Println("URL:>", "https://api.box.com/oauth2/token")
	}

	values := url.Values{}
	values.Add("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	values.Add("client_id", client_id)
	values.Add("client_secret", client_secret)
	values.Add("assertion", token)

	req, err := http.NewRequest("POST", "https://api.box.com/oauth2/token", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	fatal(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fatal(err)
	bodyString := string(bodyBytes)
	if resp.StatusCode != 200 && debug == true {
		fmt.Println("response Status:", resp.Status)
		fmt.Println(bodyString)
		return
	}
	
	var boxToken BoxToken
	if err := json.Unmarshal([]byte(bodyString), &boxToken); err != nil {
		panic(err)
	}

	fmt.Println(bodyString)
}
