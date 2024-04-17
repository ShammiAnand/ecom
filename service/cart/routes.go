package cart

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shammianand/ecom/service/auth"
	"github.com/shammianand/ecom/types"
	"github.com/shammianand/ecom/utils"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /cart/checkout", auth.WithJWTAuth(h.handleCartCheckout, h.userStore))
	router.HandleFunc("GET /orders", auth.WithJWTAuth(h.handleGetOrders, h.userStore))

	// TODO
	// router.HandleFunc("GET /cart", auth.WithJWTAuth(h.handleGetOrders, h.userStore))
	// router.HandleFunc("POST /cart", auth.WithJWTAuth(h.handleGetOrders, h.userStore))
}

func (h *Handler) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	_, err := h.userStore.GetUserByID(userId)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}
	orders, err := h.store.GetOrders(userId)
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
		orders,
	)
}

func (h *Handler) handleCartCheckout(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	cart := types.CartCheckoutPayload{}
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}
	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(
			w,
			http.StatusBadRequest,
			errors,
		)
		return
	}
	productIds, err := getProductIDsFromCartItems(cart.Items)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	products, err := h.productStore.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}
	orderId, totalPrice, err := h.createOrder(products, cart.Items, userId)
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
		map[string]any{
			"order_id":    orderId,
			"total_price": totalPrice,
		},
	)
}
