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
		Tags      []struct {
			Name string `json:"name"`
		} `json:"tags"`
		User struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			ProfileImageURL string `json:"profile_image_url"`
		} `json:"user"`
		Comments []interface{} `json:"comments"`
		Groups   []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"groups"`
	} `json:"posts"`
	Meta struct {
		PreviousPage interface{} `json:"previous_page"`
		NextPage     string      `json:"next_page"`
		Total        int         `json:"total"`
	} `json:"meta"`
}
