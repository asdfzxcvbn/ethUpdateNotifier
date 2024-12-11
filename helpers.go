package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLatestVersionForID(appid string) (*APIResult, error) {
	uuid, _ := uuid.NewRandom() // prevent cached response
	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%s&request=%v", appid, uuid)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var r APIResp
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if len(r.Results) == 0 {
		return nil, errors.New("no results found")
	}

	return &r.Results[0], nil
}
