package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	var err error
	tpl, err = template.ParseGlob("*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
		return
	}
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/tulemus", resultHandler)
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		log.Fatal("Could not start server")
		return
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "midaaaaaaaaaa")
		return
	}
	tpl.ExecuteTemplate(w, "mainpage.html", nil)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "404 Page not found")
	}
	r.FormValue()
}
