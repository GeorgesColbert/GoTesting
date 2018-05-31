package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type beauty struct{
	Title string
	Writer string
}


func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func beautify(w http.ResponseWriter, r *http.Request){
	p := beauty{Title: " Lytic test", Writer: " Georges"}
	t,_ := template.ParseFiles("basictemplating.html")
	t.Execute(w,p)
}

func beautiful(w http.ResponseWriter, r *http.Request){
	p := beauty{Title: " Lytic test", Writer: " Georges"}
	t,_ := template.ParseFiles("designtemplating.html")
	t.Execute(w,p)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/agg/", beautify)
	http.HandleFunc("/ban/", beautiful)
    http.ListenAndServe(":8080", nil)

}


