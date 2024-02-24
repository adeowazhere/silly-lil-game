package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func startPage(w http.ResponseWriter, r *http.Request) {
	var fileName = "start.html"
	page, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = page.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/start":
		startPage(w, r)
	default:
		var fileName = "index.html"
		page, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		err = page.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
