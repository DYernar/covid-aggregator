package main

import(
	"fmt"
	"net/http"
	"html/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		if r.Method == "GET" {
			news, err := GetAllNews()
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
				return
			}
			t, terr:= template.ParseFiles("templates/index.html")
			if terr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w,terr.Error())
				return
			}
			t.Execute(w, news)
		} else {
			w.WriteHeader(400)
			fmt.Fprintf(w, "<h1>400 :(<h1><br><h3>Bad request!</h3>")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "<h1>404 ;(<h1><br><h3>Page not found!</h3>")
	}
}