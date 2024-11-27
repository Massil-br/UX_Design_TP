package main

import (
	"UX_Design_TP/class"
	"UX_Design_TP/src"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type Server struct {
	Posts []class.Post
}

var server = Server{}

func Home(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index", server)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit , err := strconv.Atoi(limitStr)

	if err != nil || limit <1 {
		limit = 10
	}

	posts, err := src.GetPosts(page, limit)

	if err != nil{
		http.Error(w, "Failed to fetch posts : " + err.Error(), http.StatusInternalServerError)
		return 
	}

	server.Posts = posts

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(posts)
	if err != nil{
		http.Error(w, "Failed to encode posts to JSON: "+ err.Error(), http.StatusInternalServerError)
	}

}
