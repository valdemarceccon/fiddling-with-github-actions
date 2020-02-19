package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<h1>Olá</h1>"))

		if err != nil {
			log.Println(err.Error())
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

