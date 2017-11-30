package main

import (

	"log"
  "fmt"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	_, err := encodeJSONToGo()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

  // Postsを分解して子項目を表示する
  for _, post := range articles.Posts {
      // サンプル
      title := post.Title
      fmt.Println(title)
  }

}
