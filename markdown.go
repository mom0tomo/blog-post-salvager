package main

import (
  "bufio"
  "os"
)

func markdown(articles Article) (error) {
  // Postsを分解して子項目を表示する
  for _, post := range articles.Posts {


      title := post.Title+".md"
      body  := post.Body

      file, err := os.OpenFile("./md/"+title, os.O_WRONLY|os.O_CREATE, 0666)
      if err != nil {
        return err
      }
      defer file.Close()
      writer := bufio.NewWriter(file)
      bw := bufio.NewWriter(writer)
      bw.WriteString(body)
      bw.Flush()
  }
  return nil
}
