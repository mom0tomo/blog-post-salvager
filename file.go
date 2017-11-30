package main

import (
	"log"
	"os"
)

var title string
var body string
	
func createFiles(){
	title, body, err := getContents()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

	filename := title + ".md"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
	defer file.Close()

	file.Write(([]byte)(body))
	
	return
}

func getContents()(string, string, error) {
	articles, err := encodeJSONToGo()
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}

	for _, post := range articles.Posts {
	    title = post.Title
	    body  = post.Body
    }
    return title, body, err
}