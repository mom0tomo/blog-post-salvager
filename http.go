package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func newRequest() (*http.Request, error) {

	team_domain := os.Getenv("TEAM_DOMAIN")
	pages := os.Getenv("PAGES")
	author_id := os.Getenv("AUTHOR_ID")

	url := "https://api.docbase.io/teams/" + team_domain + "/posts?page=" + pages + "&per_page=20&q=author_id:" + author_id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// リクエストヘッダを確認する
	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	if err != nil {
		log.Fatal("Error requesting dump")
	}

	return req, err
}

func getResponse() (*http.Response, error) {
	req, err := newRequest()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}

	return resp, err
}
