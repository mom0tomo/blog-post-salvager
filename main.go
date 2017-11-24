package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  )

func main() {
  req, err := http.NewRequest("GET", "https://api.docbase.io/teams/dip-dev/posts?q=author_id:21", nil)
  if err != nil {
    fmt.Println(err)
  }
  req.Header.Set("X-Docbasetoken", "")

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
  	fmt.Println(err)
  }
  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
  defer resp.Body.Close()
}
