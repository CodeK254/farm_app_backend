package user

import (
	"fmt"
	"net/http"

	"github.com/CodeK254/farm_app_backend/cmd/services/auth"
	"github.com/CodeK254/farm_app_backend/config"
	"github.com/CodeK254/farm_app_backend/types"
	"github.com/CodeK254/farm_app_backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	/// Get JSON Payload
	var payload types.LoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	/// Check if user already exists
	user, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(payload.Password)){
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid username or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	/// Get JSON Payload
	var payload types.RegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	/// Check if user already exists
	_, err := h.store.GetUserByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}