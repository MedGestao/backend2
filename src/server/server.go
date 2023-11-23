package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type message struct {
	message string `json:"message"`
	name    string `json:"name"`
}

func OpenServer() {
	fs := http.FileServer(http.Dir("frontend/"))
	http.Handle("/", fs)

	http.HandleFunc("/api/data", handleAPIResquest)

	fmt.Println("Servidor rodando na porta: 8000!")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func handleAPIResquest(w http.ResponseWriter, r *http.Request) {
	//response := []byte(`{"Hello from the backend!"}`)
	if r.Method == "POST" {
		n := r.FormValue("name")

		response := message{
			message: "olá" + n + "!",
			name:    n,
		}

		// Imprimir a mensagem no console do servidor
		fmt.Println("Mensagem recebida:", response.message)
		fmt.Println("Nome recebido:", response.name)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
