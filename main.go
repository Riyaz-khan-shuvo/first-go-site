package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Hello I am working")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/features", featurePage)
	http.HandleFunc("/docs", docsPage)
	// to add file server

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8888", nil)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello I am from homepage ")
	pToTem, err := template.ParseFiles("./Template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}
func featurePage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./Template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wPage/feature.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	pToTem.Execute(w, nil)
}
func docsPage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./Template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	pToTem, err = pToTem.ParseFiles("./wPage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	pToTem.Execute(w, nil)
}
