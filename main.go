package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var articles Articles

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	//fmt.Fprintf(w, "ID: "+id)
	for _, item := range articles {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Println("Endpoint Hit: returnSingleArticle")
	json.NewEncoder(w).Encode(&Article{})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My Sumit's Page!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := mux.NewRouter().StrictSlash(true)
	server.HandleFunc("/", homePage)
	server.HandleFunc("/articles", returnAllArticles)
	server.HandleFunc("/article/{id}", returnSingleArticle)
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

func main() {
	articles = append(articles, Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	articles = append(articles, Article{Id: "2", Title: "Hello 2", Desc: "Article Description2", Content: "Article Content2"})
	handleRequests()
}
