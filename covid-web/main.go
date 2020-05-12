package main

import(
	"os"
	"net/http"
	"log"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}
	http.HandleFunc("/", MainPage)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}

}