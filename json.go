package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var articles Article

func getJSON() ([]byte, error) {
	resp, err := getResponse()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return byteArray, err
}

func encodeJSONToGo() (Article, error) {

	byteArray, err := getJSON()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

	if err := json.Unmarshal(byteArray, &articles); err != nil {
		log.Fatalf("Error!: %v", err)
	}
	return articles, err
}
