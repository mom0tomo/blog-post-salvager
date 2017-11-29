package main

import "time"

// https://help.docbase.io/posts/92984
type Article struct {
	Posts []struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		Draft     bool      `json:"draft"`
		URL       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
		Scope     string    `json:"scope"`
		Tags      interface{} `json:"tags"`
		User      interface{} `json:"user"`
		Comments  []interface{} `json:"comments"`
		Groups    []interface{} `json:"groups"`
	} `json:"posts"`
	Meta interface{} `json:"meta"`
}
