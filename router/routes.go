package router

import (
	"net/http"

	. "go-scrawling/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"GetPosts",
		"GET",
		"/posts",
		GetPosts,
	},
	Route{
		"GetPost",
		"GET",
		"/post/{id}",
		GetPost,
	},
	Route{
		"CreatePost",
		"POST",
		"/post/{id}",
		CreatePost,
	},
	Route{
		"DeletePost",
		"DELETE",
		"/post/{id}",
		CreatePost,
	},
}
