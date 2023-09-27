package handler

import (
	"encoding/json"
	"github.com/MrRytis/go-weather/internal/model"
	"github.com/MrRytis/go-weather/internal/service"
	"github.com/MrRytis/go-weather/internal/storage"
	"github.com/MrRytis/go-weather/pkg/response"
	"net/http"
	"time"
)

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req model.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.ErrorJSON(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	pass, err := service.HashPassword(req.Password)
	if err != nil {
		response.ErrorJSON(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	u := storage.User{
		Email:     req.Email,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.Repository.CreateUser(u)
	if err != nil {
		response.ErrorJSON(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	res := model.RegisterResponse{
		Message: "User created",
	}

	response.JSON(w, http.StatusCreated, res)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req model.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.ErrorJSON(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	u, err := h.Repository.GetUserByEmail(req.Email)
	if err != nil {
		response.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if service.CheckUserPassword(req.Password, u.Password) != nil {
		response.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	u.UpdatedAt = time.Now()
	h.Repository.UpdateUser(&u)

	token, err := service.GenerateJWT(&u)
	if err != nil {
		response.ErrorJSON(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	res := model.LoginResponse{
		Token:   token,
		Expires: time.Now().Add(service.AccessTokenJwtExpDuration),
	}

	response.JSON(w, http.StatusOK, res)
}
