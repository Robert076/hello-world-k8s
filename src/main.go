package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode("Hello World!"); err != nil {
			log.Fatal("Error writing out the message: ", err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
