package main

import (
	functions "Tsenter/Functions"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	var err error
	tpl, err = template.ParseGlob("*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	http.HandleFunc("/", mainPage)
	//see on heroku pordi jaoks. Kui on tühi siis 8080 ja kui ei ole siis launchib pordi. Üldiselt ei ole tühi
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	fmt.Println("Listening on port: ", port)
	//url kutsub ja resultHandler funktsioon käivitub
	http.HandleFunc("/furnituur", resultHandler)
	//+port võtame heroku jaoks pordi.
	err2 := http.ListenAndServe(":"+port, nil)
	if err2 != nil {
		log.Fatal("Could not start server", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "404 Page not found")
		return
	}
	tpl.ExecuteTemplate(w, "mainpage.html", nil)
}

// r method on posti meetod.
func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "404 Page not found")
	}

	var konstruktsioon []string
	var tehnoloogia []string
	// liidame arraysse valikud mis me valime veebilehel. Lisab kõik arraysse. Saame formi value.
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
	/* fmt.Println("\nKonstruktsiooni skoor") */

	// slaicib esimesed viis jätab alles.
	list = list[:6]
	//fmt.Print(list)
	// teeb struct arrayst jsoni
	serializedData, err := json.Marshal(list)
	if err != nil {
		fmt.Print("ERROR")
	}

	/* fmt.Println(string(serializedData)) */
	// saad õelda mis headerit on vaja. Saadame välja jsonit.
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON data to the response
	w.Write(serializedData)
}
