package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type NewsAggPage struct {
	Title string
	News  string
}

func page(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "yes")

}

func main() {
	log.Println("server on port 8080")
	http.HandleFunc("/", page)
	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "8080"
	}

	log.Println(http.ListenAndServe(":"+port, nil))

}
