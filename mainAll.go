package main

import (
  "fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
  "encoding/json"
  "io/ioutil"
  "bufio"
  "strings"
  "time"

	"github.com/joho/godotenv"
)

var articles Article

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
		User struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			ProfileImageURL string `json:"profile_image_url"`
		} `json:"user"`
		Comments []interface{} `json:"comments"`
		Groups   []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"groups"`
	} `json:"posts"`
	Meta struct {
		PreviousPage interface{} `json:"previous_page"`
		NextPage     string      `json:"next_page"`
		Total        int         `json:"total"`
	} `json:"meta"`
}

func init() {
	_ = godotenv.Load()
}


func getError(err error) {
  log.Fatalf("Error!: %v", err)
}


func main() {
  var url string
  // URLの生成
  url = "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?&q=author_id:" + os.Getenv("AUTHOR_ID")

  // リクエストの生成
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    getError(err)
  }
  req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
  req.Header.Set("Content-Type", "application/json")

  // リクエストヘッダを確認する
  dump, err := httputil.DumpRequestOut(req, true)
  fmt.Printf("%s", dump)
  if err != nil {
    log.Fatal("Error requesting dump")
  }

  // 疎通確認をしリクエストを受け取る
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    getError(err)
  } else if res.StatusCode != 200 {
    getError(fmt.Errorf("Unable to get this url : http status %d", res.StatusCode))
  }

  // バイトデータとして読み込む
  byteArray, err := ioutil.ReadAll(res.Body)
  if err != nil {
    getError(err)
  }
  defer res.Body.Close()

  // Jsonの形を構造体へ突っ込む
  if err := json.Unmarshal(byteArray, &articles); err != nil {
    getError(err)
  }

  // Articleの子であるPostsを親にしそれぞれのデータを取得できる状態にする。(PHPオブジェクトと同じ)
  for _, post := range articles.Posts {
    title := strings.Replace(post.Title, "/", "-", -1)
    body  := post.Body
    func() {
        file, err := os.OpenFile("./md/"+title+".md", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
          getError(err)
        }
        defer file.Close()
        writer := bufio.NewWriter(file)
        bw := bufio.NewWriter(writer)
        bw.WriteString(body)
        bw.Flush()
    }()
    fmt.Println(title)
  }

  // Meta情報。　ここが欲しかっただけ
  url = articles.Meta.NextPage
  total := articles.Meta.Total
  cntLoop := (total/20)
  if total % 20 >= 1 {
    cntLoop += 1
  }


  // ループに必要な材料は揃ったので-----------Loop Time
  for i := 0; i < cntLoop; i++ {

    // next_pageを元URLとして扱うことにする
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      getError(err)
    }
    req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
    req.Header.Set("Content-Type", "application/json")

    // 疎通確認をしリクエストを受け取る
    res, err := http.DefaultClient.Do(req)
    if err != nil {
      getError(err)
    } else if res.StatusCode != 200 {
      getError(fmt.Errorf("Unable to get this url : http status %d", res.StatusCode))
    }

    // バイトデータとして読み込む
    byteArray, err := ioutil.ReadAll(res.Body)
    if err != nil {
      getError(err)
    }
    defer res.Body.Close()

    // Jsonの形を構造体へ突っ込む
    if err := json.Unmarshal(byteArray, &articles); err != nil {
      getError(err)
    }

    // Articleの子であるPostsを親にしそれぞれのデータを取得できる状態にする。(PHPオブジェクトと同じ)
    for _, post := range articles.Posts {
      title := strings.Replace(post.Title, "/", "-", -1)
      body  := post.Body
      func() {
          file, err := os.OpenFile("./md/"+title+".md", os.O_WRONLY|os.O_CREATE, 0666)
          if err != nil {
            getError(err)
          }
          defer file.Close()
          writer := bufio.NewWriter(file)
          bw := bufio.NewWriter(writer)
          bw.WriteString(body)
          bw.Flush()
      }()
      fmt.Println(title)
    }
    // next_pageを取り続ける
    url = articles.Meta.NextPage
  }
}
