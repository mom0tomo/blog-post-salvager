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

  for _, post := range articles.Posts {
      fmt.Printf("Title: %v\n", post.Body)
  }
}
