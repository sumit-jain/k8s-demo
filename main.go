package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/", homePage)
	server.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

func main() {
	handleRequests()
}
