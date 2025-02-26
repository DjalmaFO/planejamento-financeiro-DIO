package adapter_http

import (
	"fmt"
	"net/http"

	transaction_handler "github.com/DjalmaFO/planejamento-financeiro-DIO/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/transactions", transaction_handler.GetTransactions)
	http.HandleFunc("/transaction/create", transaction_handler.CreateTransaction)

	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Alive\n")
}
