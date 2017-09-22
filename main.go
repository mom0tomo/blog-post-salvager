package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"net/http/httputil"
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
	url := "https://dip-dev.docbase.io/posts/289087"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-DocBaseToken", "F-xs-QV5xpekgU5Zu8xj")
	req.Header.Set("Content-Type", "appliation/json")
	if err != nil {
		return nil, err
	}
	// リクエストヘッダの内容を確認する
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	// DocBaseのAPIを叩く
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

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
	// 記事を表示する
	fmt.Println(articles)
}
