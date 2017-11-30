package main

import (
	"reflect"
	"log"
	"os"
	"fmt"
	"path/filepath"	
)

func createFiles(){
	title := "title"
	body  := "body"
	// title, body, err := getContents()
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

	// articles構造体の中のpost構造体のtitle/bodyフィールドを取得する
	t := reflect.ValueOf(articles)
	p := reflect.ValueOf(t.Field(0))
	title := p.Field(1).Name
	// body  := p.Field(2).Name
	fmt.Println(title)
	return title, body, err
	// return title, body, err
}