package main

import (
	functions "Tsenter/Functions"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	var err error
	tpl, err = template.ParseGlob("*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
		return
	}
	http.HandleFunc("/", mainPage)

	http.HandleFunc("/furnituur", resultHandler)
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

	var konstruktsioon []string
	var tehnoloogia []string

	konstruktsioon = append(konstruktsioon, r.FormValue("k1"))
	konstruktsioon = append(konstruktsioon, r.FormValue("k2"))
	konstruktsioon = append(konstruktsioon, r.FormValue("k3"))
	konstruktsioon = append(konstruktsioon, r.FormValue("mont/fp"))
	konstruktsioon = append(konstruktsioon, r.FormValue("nahtavus1"))
	konstruktsioon = append(konstruktsioon, r.FormValue("nahtavus2"))

	tehnoloogia = append(tehnoloogia, r.FormValue("CNC"))
	tehnoloogia = append(tehnoloogia, r.FormValue("servapuurimisevoimekus"))
	tehnoloogia = append(tehnoloogia, r.FormValue("Lamello"))
	tehnoloogia = append(tehnoloogia, r.FormValue("Domino"))
	tehnoloogia = append(tehnoloogia, r.FormValue("OVVO"))

	list := functions.ScoreCalc(konstruktsioon, tehnoloogia)
	fmt.Println("\nKonstruktsiooni skoor")

	list = list[:6]
	//fmt.Print(list)
	serializedData, err := json.Marshal(list)
	if err != nil {
		fmt.Print("ERROR")
		return
	}

	fmt.Println(string(serializedData))
}
