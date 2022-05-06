package main

import (
	"encoding/json"
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"io/ioutil"
	"time"
)

const (
	baseURL = "https://rickandmortyapi.com/api"
	timeout = 3000 * time.Millisecond
)

type CharacterResponse struct {
	Info    Metadata
	Results []Character
}

type Metadata struct {
	Count int
	Pages int
	Next  string
	Prev  string
}

type Character struct {
	ID      int
	Name    string
	Status  string
	Species string
	Type    string
	Gender  string
}

func main() {
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get(baseURL+"/character", nil)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	var response CharacterResponse
	json.Unmarshal(body, &response)
	fmt.Printf("API Response %+v", response)
}
