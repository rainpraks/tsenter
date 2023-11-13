package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", mainPage)

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "lala")
}
