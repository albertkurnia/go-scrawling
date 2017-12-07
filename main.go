package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	. "go-scrawling/data"
	. "go-scrawling/handlers"
)

func main() {

	//create connection with target
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
		Posts = append(Posts, Post{ID: strconv.Itoa(index), Title: title, Link: link})
	})

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")

	fmt.Print("Listen to 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
