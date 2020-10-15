package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var adminTemplateAbsPath string
var workingDirectory string
var websiteTemplateAbsPath string

func init() {

	workingDirectory, _ = os.Getwd()
	//fmt.Println(workingDirectory)

	websiteTemplateAbsPath = filepath.Join(workingDirectory, "templates", "website", "*.gohtml")
	//adminTemplateAbsPath = filepath.Join(workingDirectory, "templates", "admin")
	//fmt.Println(websiteTemplateAbsPath)
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/prospectus", prospectus)
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

func index(w http.ResponseWriter, r *http.Request) {

	//order confirm
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("catch panic >>", r)
		}
	}()

	//fmt.Println(">>", websiteTemplateAbsPath)
	//ptmp, err := template.New("webtemplate.gohtml").ParseGlob(websiteTemplateAbsPath)
	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(ptmp)

	//pageName := filepath.Join(workingDirectory, "page", "dashboard.gohtml")
	// ptmp, err = ptmp.ParseFiles(pageName)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// data := struct {
	// 	Title string
	// }{
	// 	Title: "Homepage",
	// }

	ptmp.Execute(w, nil)

}

func prospectus(w http.ResponseWriter, r *http.Request) {

	//order confirm
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("catch panic >>", r)
		}
	}()

	//fmt.Println(">>", websiteTemplateAbsPath)
	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(ptmp)

	pageName := filepath.Join(workingDirectory, "page", "prospectus.gohtml")
	ptmp, err = ptmp.ParseFiles(pageName)
	if err != nil {
		fmt.Println(err.Error())
	}

	// data := struct {
	// 	Title string
	// }{
	// 	Title: "Homepage",
	// }

	ptmp.Execute(w, nil)

}
