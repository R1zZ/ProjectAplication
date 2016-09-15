package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request){
fmt.Fprintln(w,"<h1>Hello, world!</h1>")
}

func NameHandler(w http.ResponseWriter, r *http.Request)  {
  hallo := r.FormValue("hallo")
  nama := r.FormValue("nama aing")
  fmt.Fprintln(w,"<h1>%s, %s!</h1>", hallo, nama)
}

func main(){
http.HandleFunc("/", RootHandler)
http.HandleFunc("/nama", NameHandler)
http.ListenAndServe(":8000",nil)
}
