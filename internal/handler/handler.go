package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/usecase"
)

// Handler contains the handlers to HTTP API.
type Handler struct {
	usecase usecase.CustomerUsecase
}

// PostCustomer handle a new request to connect with create customer usecase.
func (h Handler) PostCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.usecase.Create(r.Context(), customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// New returns a instance of Handler.
func New(usecase usecase.CustomerUsecase) Handler {
	return Handler{
		usecase: usecase,
	}
}
