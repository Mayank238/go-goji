package main

import (
	//"fmt"
	"net/http"
  "html/template"
	"goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//_ = pat.Param(r, "name")
	//fmt.Fprintf(w, "Hello!")
	t ,err := template.ParseFiles("demo.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w,nil)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), hello)

	http.ListenAndServe("localhost:8000", mux)
}
