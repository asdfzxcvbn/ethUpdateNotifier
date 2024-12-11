package main

type APIResult struct {
	Name    string `json:"trackName"`
	Version string `json:"version"`
}

type APIResp struct {
	Results []APIResult `json:"results"`
}
