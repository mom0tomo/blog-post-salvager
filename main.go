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
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Draft     bool      `json:"draft"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []struct {
		Name string `json:"name"`
	} `json:"tags"`
	Scope  string `json:"scope"`
	Groups []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	User struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		ProfileImageURL string `json:"profile_image_url"`
	} `json:"user"`
}

// 環境変数を取得する
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func fetch() ([]Article, error) {
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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}
	// JSONを読み込む
	body, err := ioutil.ReadAll(resp.Body)

	// JSONを構造体へデコードする
	var articles []Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func main() {
	articles, err := fetch()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
	fmt.Println(articles)
}
