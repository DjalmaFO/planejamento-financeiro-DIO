package transaction_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DjalmaFO/planejamento-financeiro-DIO/model/transaction"
	"github.com/DjalmaFO/planejamento-financeiro-DIO/util"
)

const (
	contentType      = "Content-Type"
	valueContentType = "application/json"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set(contentType, valueContentType)

	createdAt, _ := util.StringToTime(time.Now().Format(util.LayoutDate), util.LayoutDate)

	transactions := []transaction.Transacition{
		{
			Title:     "Salário",
			Amount:    1200.00,
			Type:      0,
			CreatedAt: createdAt,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var transaction transaction.Transacition

	bts, _ := io.ReadAll(r.Body)

	json.Unmarshal(bts, &transaction)

	fmt.Println(transaction)
}
