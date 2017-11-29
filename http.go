package main

import (
  "fmt"
	"log"
  "os"
	"net/http"
	"net/http/httputil"
)

func newRequest()(*http.Request, error) {
  url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?q=author_id:" + os.Getenv("AUTHOR_ID")

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

func getResponse()(*http.Response, error) {
  req, err := newRequest()

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  } else if resp.StatusCode != 200 {
    return nil, fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
  }

  return resp, err
}