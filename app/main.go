package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	payload := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("erro ao serializar resposta: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", projetoKorpHandler)

	log.Println("http-server-projeto-korp iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
