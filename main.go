package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit!")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Desc:"An example article", Content:"Hello world!"},
	} 

	fmt.Fprintf(w, "All Articles endpoint hit!")
	json.NewEncoder(w).Encode(articles)
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Article endpoint hit!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/", homePage).Methods("GET")
	http.HandleFunc("/articles", allArticles).Methods("GET")
	http.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}