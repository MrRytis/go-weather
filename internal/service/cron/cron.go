package cron

import (
	"fmt"
	"github.com/MrRytis/go-weather/internal/service/cron/jobs"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Job interface {
	GetName() string
	RunJob(*gorm.DB)
}

func IsJobAvailable(name string) bool {
	cronJobs := getJobs()

	for _, job := range cronJobs {
		if job.GetName() == name {
			return true
		}
	}

	return false
}

func AddJob(c *cron.Cron, db *gorm.DB, name string, spec string) error {
	job, err := GetJob(name)
	if err != nil {
		return err
	}

	if spec == "" {
		spec = "@every 3h"
	}

	c.Entries()

	_, err = c.AddFunc(spec, func() {
		job.RunJob(db)
	})
	if err != nil {
		return err
	}

	return nil
}

func GetJob(name string) (Job, error) {
	cronJobs := getJobs()

	for _, job := range cronJobs {
		if job.GetName() == name {
			return job, nil
		}
	}

	return nil, fmt.Errorf("job %s is not available", name)
}

func getJobs() []Job {
	return []Job{
		&jobs.WeatherCronJob{},
	}
}
