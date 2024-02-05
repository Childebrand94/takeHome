package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/logic"
	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
)

type Query struct{}

func (q *Query) Parse(w http.ResponseWriter, r *http.Request) {
	_, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	var payload models.Query

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}
	defer r.Body.Close()
	data, err := logic.ProcessData(payload)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Error processing data", err)
	}

	resp, err := json.Marshal(data)
	if err != nil {
		models.SendError(w, http.StatusInternalServerError, "Failed to prepare response", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
