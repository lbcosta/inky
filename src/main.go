package main

import (
	"inky/src/modules/db"
	"inky/src/modules/gmail"
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
				gmail.PrintUserLabels()
				w.WriteHeader(200)
				w.Write([]byte("Printed successfully"))
			}
		})

	http.DefaultServeMux.
		HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				w.WriteHeader(200)
				w.Write([]byte("You can now return to the application."))
			}
		})

	if err := http.ListenAndServe("localhost:3456", nil); err != nil {
		panic(err)
	}
}
