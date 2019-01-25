package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getMessage(url string) (*Message, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Panic(err)
	}

	return &msg, err
}

func getMessages(url string) (*Messages, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var msgs Messages
	err = json.Unmarshal(body, &msgs)
	if err != nil {
		log.Panic(err)
	}

	return &msgs, err
}