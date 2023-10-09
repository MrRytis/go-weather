package model

type CronJobRequest struct {
	JobName string `json:"job" validate:"required" example:"weather"`
	Spec    string `json:"spec" example:"@every 3h"`
}

type CronResponse struct {
	Message string `json:"message"`
}
