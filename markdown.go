package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"strings"
)

func closeFile(file *os.File) {
	file.Close()
}

func createFile(title string) (*os.File, error) {
  // [WIP]./md/はディレクトリを入力させるか、envで指定してあげると良さそう
	file, err := os.OpenFile("./md/"+title+".md", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func markdown(articles Article) (error) {


	for _, post := range articles.Posts {
		title := strings.Replace(post.Title, "/", "-", -1)
		body  := post.Body
		func() {
				file, err := createFile(title)
				if err != nil {
					log.Fatalf("Error!: %v", err)
				}
				defer closeFile(file)
				writer := bufio.NewWriter(file)
				bw := bufio.NewWriter(writer)
				bw.WriteString(body)
				bw.Flush()
		}()
	}
	fmt.Println("終了")
	return nil
}
