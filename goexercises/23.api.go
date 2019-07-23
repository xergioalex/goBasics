package main

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main () {
	var client = &http.Client{ Timeout: 10 * time.Second }
	url := "http://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)

	if err != nil {
		panic(err.Error)
	}

	var post [] Post

	err = json.NewDecoder(resp.Body).Decode(&post)

	if err != nil {
		panic(err.Error)
	}

	fmt.Println(len(post))
	fmt.Println(post[0])
}