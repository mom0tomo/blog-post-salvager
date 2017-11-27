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
		User     []interface{} `json:"user"`
		Comments []interface{} `json:"comments"`
		Groups   []interface{} `json:"groups"`
	} `json:"posts"`
	Meta []interface{} `json:"meta"`
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func fetch() error {
	loadEnv()

	url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?q=author_id:" + os.Getenv("AUTHOR_ID")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
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
		return err
	} else if resp.StatusCode != 200 {
		return fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// JSONを読み込む
	body, _ := ioutil.ReadAll(resp.Body)

	articles := new(Article)
	err = json.Unmarshal(body, articles)

	fmt.Println(articles)
	return err
}

func main() {
	err := fetch()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
}
