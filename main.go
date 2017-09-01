package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func fetch() ([]Article, error) {
	// DocBaseのAPIを叩く
	res, err := http.Get("https://dip-dev.docbase.io/posts/280557")
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// JSONを読み込む
	body, err := ioutil.ReadAll(res.Body)

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
	// 記事を表示する
	fmt.Println(articles)
}
