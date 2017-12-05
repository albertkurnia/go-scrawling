package main

type Post struct {
	ID    string `json:"id, omitempty"`
	Title string `json:"title, omitempty"`
	Link  string `json:"_link, omitempty"`
}
