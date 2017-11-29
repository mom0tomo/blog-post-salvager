package main

import (
	"fmt"
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

	fmt.Println(articles)
}
