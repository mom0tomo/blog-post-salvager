package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Json to Goを使ってAPIで取得できる情報を構造体に入れる
type Article struct {
	Posts []struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		Draft     bool      `json:"draft"`
		URL       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
		Scope     string    `json:"scope"`
		Tags      []struct {
			Name string `json:"name"`
		} `json:"tags"`
		User     interface{} `json:"user"`
		Comments []interface{} `json:"comments"`
		Groups   []interface{} `json:"groups"`
	} `json:"posts"`
	Meta interface{} `json:"meta"`
}

func main() {
	articles, err := unmarshalJSON()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
  fmt.Println(articles)
}

func loadEnv() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func fetchJson()([]byte, error) {
  loadEnv()

  url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?q=author_id:" + os.Getenv("AUTHOR_ID")

  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return nil, err
  }

  // リクエストヘッダの内容を確認する
  dump, err := httputil.DumpRequestOut(req, true)
  fmt.Printf("%s", dump)
  if err != nil {
    log.Fatal("Error requesting dump")
  }

  // DocBaseのAPIを叩く
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  } else if resp.StatusCode != 200 {
    return nil, fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
  }
  defer resp.Body.Close()

  // JSONを読み込む
  resBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  return resBody, nil
}

func unmarshalJSON()(Article, error) {
  var articles Article

  resBody, err := fetchJson()
  if err != nil {
    log.Fatalf("Error!: %v", err)
  }

  err = json.Unmarshal(resBody, &articles)
  if err != nil {
    log.Fatalf("Error!: %v", err)
  }
  return articles, err
}
