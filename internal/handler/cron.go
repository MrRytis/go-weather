package handler

import (
	"github.com/MrRytis/go-weather/internal/model"
	cronService "github.com/MrRytis/go-weather/internal/service/cron"
	"github.com/MrRytis/go-weather/pkg/httpUtils"
	"net/http"
)

// StartCronHandler godoc
// @Router /cron/start [post]
// @Summary Start cron jobs
// @Description Start cron jobs
// @Tags cron
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.CronResponse "Cron started"
func (h *Handler) StartCronHandler(w http.ResponseWriter, r *http.Request) {
	h.Cron.Start()

	res := model.CronResponse{
		Message: "Cron started",
	}

	httpUtils.JSON(w, http.StatusOK, res)
}

// StopCronHandler godoc
// @Router /cron/stop [post]
// @Summary Stop cron jobs
// @Description Stop cron jobs
// @Tags cron
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.CronResponse "Cron stopped"
func (h *Handler) StopCronHandler(w http.ResponseWriter, r *http.Request) {
	h.Cron.Stop()

	res := model.CronResponse{
		Message: "Cron stopped",
	}

	httpUtils.JSON(w, http.StatusOK, res)
}

// AddCronJobHandler godoc
// @Router /cron/add [put]
// @Summary Add cron job
// @Description Add cron job
// @Tags cron
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body model.CronJobRequest true "Cron job details"
// @Success 200 {object} model.CronResponse "Cron added"
// @Failure 400 {object} httpUtils.ErrorResponse "Invalid request body"
// @Failure 500 {object} httpUtils.ErrorResponse "Failed to add job"
func (h *Handler) AddCronJobHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CronJobRequest
	err := httpUtils.ParseJSON(r, w, &req)
	if err != nil {
		return
	}

	err = httpUtils.ValidateStruct(w, req)
	if err != nil {
		return
	}

	if !cronService.IsJobAvailable(req.JobName) {
		httpUtils.ErrorJSON(w, "Job is not available", http.StatusBadRequest)
		return
	}

	if req.Spec == "" {
		req.Spec = "@every 3h"
	}

	err = cronService.AddJob(h.Cron, h.Repository.Db, req.JobName, req.Spec)
	if err != nil {
		httpUtils.ErrorJSON(w, "Failed to add job", http.StatusInternalServerError)
	}

	res := model.CronResponse{
		Message: "Cron added",
	}

	httpUtils.JSON(w, http.StatusOK, res)
}
