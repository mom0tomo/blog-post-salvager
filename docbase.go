package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"


	"github.com/joho/godotenv"
)

var articles Article

func init() {
	_ = godotenv.Load()
}

func main() {
	// firstAction 
	url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?page=" + os.Getenv("PAGES") + "&per_page=20&q=author_id:" + os.Getenv("AUTHOR_ID")
	req := newRequest(url)
	resp := getResponse(req)
	jsonToStruct(resp)

	for articles.Meta.NextPage != "" {
		url = articles.Meta.NextPage
		req = newRequest(url)
		resp = getResponse(req)
		jsonToStruct(resp)
		fmt.Println(url)
	}
	

}

func newRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("(newRequest)Error!: %v", err)
	}
	req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func getResponse(req *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("(getResponse)Error!: %v", err)
	} else if resp.StatusCode != 200 {
		fmt.Errorf("Unable to get this url : http status %d", resp.StatusCode)
	}
	return resp
}

func jsonToStruct(resp *http.Response) {
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("(jsonToStruct)Error!: %v", err)
	}
	defer resp.Body.Close()

	err = json.Unmarshal(byteArray, &articles)
	if err != nil {
		log.Fatalf("(jsonToStruct)Error!: %v", err)
	}
	return
}

// func closeFile(file *os.File) {
// 	file.Close()
// }

// 
// func createMarkdown() {
// 	for _, post := range articles.Posts {
// 		title := strings.Replace(post.Title, "/", "-", -1)
// 		body := post.Body
// 		func() {
// 			file, err := os.OpenFile(os.Getenv("SAVE_DIR")+title+".md", os.O_WRONLY|os.O_CREATE, 0666)
// 			if err != nil {
// 				log.Fatalf("Error!: %v", err)
// 			}
// 			defer closeFile(file)
// 			writer := bufio.NewWriter(file)
// 			bw := bufio.NewWriter(writer)
// 			bw.WriteString(body)
// 			bw.Flush()
// 		}()
// 		fmt.Println(title + ".md")
// 	}
// 	fmt.Println("終了：" + os.Getenv("SAVE_DIR") + "に保存されました")
// }