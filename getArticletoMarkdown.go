package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var articles Article

func init() {
	_ = godotenv.Load()
}

func closeFile(file *os.File) {
	file.Close()
}

func main() {
	pages := os.Getenv("PAGES")

	if (pages == 0) {
		getLimitArticles()
	} else {
		getSpecifyArticles()
	}
}


func getSpecifyArticles() {

}

func getLimitArticles() {

}

func main() {
	team_domain := os.Getenv("TEAM_DOMAIN")
	author_id := os.Getenv("AUTHOR_ID")
	save_dir := os.Getenv("SAVE_DIR")
	
	url := "https://api.docbase.io/teams/" + team_domain + "/posts?page=" + pages + "&per_page=20&q=author_id:" + author_id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

	req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// リクエストヘッダを確認する
	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	} else if resp.StatusCode != 200 {
		fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(byteArray, &articles); err != nil {
		log.Fatalf("Error!: %v", err)
	}

	for _, post := range articles.Posts {
		title := strings.Replace(post.Title, "/", "-", -1)
		body := post.Body
		func() {
			file, err := os.OpenFile(save_dir+title+".md", os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("Error!: %v", err)
			}
			defer closeFile(file)
			writer := bufio.NewWriter(file)
			bw := bufio.NewWriter(writer)
			bw.WriteString(body)
			bw.Flush()
		}()
		fmt.Println(title + ".md")
	}
	fmt.Println("終了：" + save_dir + "に保存されました")
}