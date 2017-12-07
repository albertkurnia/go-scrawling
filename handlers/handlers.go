package handlers

import (
	"encoding/json"
	"go-scrawling/data"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data.Posts); err != nil {
		panic(err)
	}

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range data.Posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&data.Post{}); err != nil {
		panic(err)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range data.Posts {
		if item.ID == params["id"] {
			data.Posts = append(data.Posts[:index], data.Posts[index+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data.Posts); err != nil {
		panic(err)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post data.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = params["id"]
	data.Posts = append(data.Posts, post)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data.Posts); err != nil {
		panic(err)
	}
}
