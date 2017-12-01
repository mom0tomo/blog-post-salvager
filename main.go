package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	articles, err := encodeJSONToGo()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

  err = markdown(articles)
  if err != nil {
    log.Fatalf("Error!: %v", err)
  }
}
