package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP recebidas por endpoint e status",
		},
		[]string{"endpoint", "method", "status"},
	)

	serviceUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_up",
			Help: "Disponibilidade do serviço: 1 = online, 0 = offline",
		},
	)
)

func init() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(serviceUp)

	serviceUp.Set(1)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		totalRequests.WithLabelValues("/projeto-korp", r.Method, "405").Inc()
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	payload := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("erro ao serializar resposta: %v", err)
		totalRequests.WithLabelValues("/projeto-korp", r.Method, "500").Inc()
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsonData); err != nil {
		log.Printf("erro ao escrever resposta: %v", err)
		return
	}

	totalRequests.WithLabelValues("/projeto-korp", r.Method, "200").Inc()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/projeto-korp", projetoKorpHandler)

	mux.Handle("/metrics", promhttp.Handler())

	log.Println("http-server-projeto-korp iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
