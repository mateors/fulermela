package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", homepage)
	http.HandleFunc("/about", aboutus)
	http.ListenAndServe(":8050", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {

	ptm, err := template.ParseFiles("./index.html")
	if err != nil {

		fmt.Println("Something wrong:", err.Error())
	}

	ptm.Execute(w, nil)
	//fmt.Fprint(w, `WELCOME TO HOMEPAGE`)
}

func aboutus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h2>Hello</h2>`)
}
