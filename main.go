package main

import (
	"log"
	"net/http"

	"github.com/eldmark/go-http/handlers"
)

func main() {

	handler := handlers.NewCharacterHandler("./data/onepiece.json")

	http.HandleFunc("/api/ping", handler.Ping)
	http.HandleFunc("/api/characters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCharacters(w, r)
		case http.MethodPost:
			handler.AddCharacter(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/characters/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCharacterByID(w, r)
		case http.MethodPut:
			handler.UpdateCharacter(w, r)
		case http.MethodDelete:
			handler.DeleteCharacter(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on :24229")
	log.Fatal(http.ListenAndServe(":24229", nil))
}
