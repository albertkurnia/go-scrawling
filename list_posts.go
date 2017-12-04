package main

import (
	// import standard libraries
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id, omitempty"`
	Title string `json:"title, omitempty"`
	Link  string `json:"_link, omitempty"`
}

var posts []Post

func postScrape() {
	doc, err := goquery.NewDocument("https://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("#main article .entry-content ul li").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		posts = append(posts, Post{ID: strconv.Itoa(index), Title: title, Link: link})
	})

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPost).Methods("GET")

	fmt.Print("Listen to 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Post{})
}

func main() {
	postScrape()
}
