package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	// "time"

	"github.com/joho/godotenv"
)

// type Article struct {
//     Posts []struct {
//         ID        int       `json:"id"`
//         Title     string    `json:"title"`
//         Body      string    `json:"body"`
//         Draft     bool      `json:"draft"`
//         URL       string    `json:"url"`
//         CreatedAt time.Time `json:"created_at"`
//         Scope     string    `json:"scope"`
//         Tags      []struct {
//             Name string `json:"name"`
//         } `json:"tags"`
//         User struct {
//             ID              int    `json:"id"`
//             Name            string `json:"name"`
//             ProfileImageURL string `json:"profile_image_url"`
//         } `json:"user"`
//         Comments []struct {
//             ID        int       `json:"id"`
//             Body      string    `json:"body"`
//             CreatedAt time.Time `json:"created_at"`
//             User      struct {
//                 ID              int    `json:"id"`
//                 Name            string `json:"name"`
//                 ProfileImageURL string `json:"profile_image_url"`
//             } `json:"user"`
//         } `json:"comments"`
//         Groups []interface{} `json:"groups"`
//     } `json:"posts"`
//     Meta struct {
//         PreviousPage interface{} `json:"previous_page"`
//         NextPage     string      `json:"next_page"`
//         Total        int         `json:"total"`
//     } `json:"meta"`
// }

var articles interface{}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func fetch()(interface{}, error) {
	loadEnv()

	url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?q=author_id:" + os.Getenv("AUTHOR_ID")
	// url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts/316251"
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
	body, err := ioutil.ReadAll(resp.Body)

	// JSONを構造体へデコードする
	// var articles []Article
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
