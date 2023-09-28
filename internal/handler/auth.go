package handler

import (
	"github.com/MrRytis/go-weather/internal/model"
	"github.com/MrRytis/go-weather/internal/service/auth"
	"github.com/MrRytis/go-weather/internal/storage"
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"net/http"
	"time"
)

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req model.AuthRequest
	err := httpUtils.ParseJSON(r, w, &req)
	if err != nil {
		return
	}

	err = httpUtils.ValidateStruct(w, req)
	if err != nil {
		return
	}

	pass, err := authService.HashPassword(req.Password)
	if err != nil {
		httpUtils.ErrorJSON(w, "Failed to hash password", http.StatusInternalServerError)
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
		httpUtils.ErrorJSON(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	res := model.RegisterResponse{
		Message: "User created",
	}

	httpUtils.JSON(w, http.StatusCreated, res)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req model.AuthRequest
	err := httpUtils.ParseJSON(r, w, &req)
	if err != nil {
		return
	}

	u, err := h.Repository.GetUserByEmail(req.Email)
	if err != nil {
		httpUtils.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if authService.CheckUserPassword(req.Password, u.Password) != nil {
		httpUtils.ErrorJSON(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	u.UpdatedAt = time.Now()
	h.Repository.UpdateUser(&u)

	token, err := authService.GenerateJWT(&u)
	if err != nil {
		httpUtils.ErrorJSON(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	res := model.LoginResponse{
		Token:   token,
		Expires: time.Now().Add(authService.AccessTokenJwtExpDuration),
	}

	httpUtils.JSON(w, http.StatusOK, res)
}
