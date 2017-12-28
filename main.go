package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	//"github.com/gorilla/mux"
	. "go-scrawling/data"
	. "go-scrawling/router"
	//. "go-scrawling/handlers"
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

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
