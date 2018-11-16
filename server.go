package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Form struct {
	Username string
	Email    string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadPage(filename string) []byte {
	body, _ := ioutil.ReadFile(filename)
	return body
}
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		f := ParseFrom(w, r)
		renderForm(w, f)
	default:
		index := loadPage("static/index.html")
		fmt.Fprint(w, string(index))
	}
}

func ParseFrom(w http.ResponseWriter, r *http.Request) Form {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error")
		return Form{"null", "null"}
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	f := Form{name, email}
	fmt.Println(f)
	return f
}

func renderForm(w http.ResponseWriter, f Form) {
	fmt.Println("render")
	t, _ := template.ParseFiles("template/form.html")
	t.Execute(w, f)
}
