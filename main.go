package main

import (
	"UX_Design_TP/src"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	src.InitDb()
	//src.SeedData(1000)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/api/posts", GetPostsHandler)

	fmt.Println("server started on port" + port)
	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
