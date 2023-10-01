package main

import (
	"articlesfeedapi/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Starting ArticlesFeedAPI")
	r := mux.NewRouter()
	r.HandleFunc("/feed", handlers.GetLatestArticlesHandler).Methods("GET")
	r.HandleFunc("/sources", handlers.GetSourcesHandler).Methods("GET")

	fmt.Println("Server started at :8084")
	http.Handle("/", r)
	http.ListenAndServe(":8084", nil)
	fmt.Println("ArticlesFeedAPI is running...")
}
