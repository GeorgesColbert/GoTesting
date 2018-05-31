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
	fmt.Fprintf(w, "Hello World i am")
	fmt.Fprintf(w, "<img src='/images/anna-bg.png' alt='gopher' style='width:235px;height:320px;'>")
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
	http.HandleFunc("/", beautiful)
	http.HandleFunc("/agg/", beautify)
	http.HandleFunc("/ban/", beautiful)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
    http.ListenAndServe(":8080", nil)

}


