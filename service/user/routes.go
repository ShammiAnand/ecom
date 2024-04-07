package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/shammianand/ecom/config"
	"github.com/shammianand/ecom/service/auth"
	"github.com/shammianand/ecom/types"
	"github.com/shammianand/ecom/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	payload := types.LoginUserPayload{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid payload %v", errors),
		)
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			errors.New("invalid email or password"),
		)
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			errors.New("incorrect password"),
		)
		return
	}

	token, err := auth.CreateJWT([]byte(config.Envs.Secret), u.ID)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	payload := types.RegisterUserPayload{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("invalid payload %v", errors),
		)
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			errors.New("user with the given email already exists"),
		)
		return
	}

	hashedPassword, err := auth.HashPasswords(payload.Password)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		Email:     payload.Email,
		LastName:  payload.LastName,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			errors.New("failed to create user"),
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "user created",
		},
	)

}
