package model

type CronJobRequest struct {
	JobName string `json:"job" validate:"required"`
	Spec    string `json:"spec"`
}

type CronResponse struct {
	Message string `json:"message"`
}
