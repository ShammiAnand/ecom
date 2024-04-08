package order

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shammianand/ecom/types"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/order", h.handleCartCheckout).Methods("POST")
}

func (h *Handler) handleCartCheckout(w http.ResponseWriter, r *http.Request) {
}
