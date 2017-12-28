package handlers

import (
	"encoding/json"
	"go-scrawling/data"
	"io"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	var post data.Post

	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = params["id"]
	data.Posts = append(data.Posts, post)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(data.Posts); err != nil {
		panic(err)
	}
}
