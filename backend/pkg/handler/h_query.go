package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/logic"
	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
)

type Query struct{}

func (q *Query) Parse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hitting pars endpoint.")
	_, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer ctxCancel()

	var payload models.Query

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		models.SendError(w, http.StatusBadRequest, "Bad request format", err)
		return
	}
	defer r.Body.Close()
	logic.ProcessData(payload)
	fmt.Printf("this is the request: %+v\n", payload)

}
