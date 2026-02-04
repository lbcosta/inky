package main

import (
	"inky/src/modules/db"
	"inky/src/modules/printer"
	"net/http"
)

func main() {
	http.DefaultServeMux.
		HandleFunc("/print", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				randomString, err := db.Get()
				if err != nil {
					w.WriteHeader(400)
					w.Write([]byte(err.Error()))
				}
				printer.Print("Inky says: " + randomString)
			}
		})

	if err := http.ListenAndServe("localhost:3456", nil); err != nil {
		panic(err)
	}
}
