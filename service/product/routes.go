package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shammianand/ecom/types"
	"github.com/shammianand/ecom/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	payload := types.Product{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}
	err := h.store.CreateProduct(payload)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "product created",
		},
	)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		products,
	)
}
