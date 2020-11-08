package service

import (
	"fmt"
	"time"
)

type SchedulerService interface {
	Start()
}

type SchedulerServiceImpl struct {
	IntervalMinutes int64
	ScheduledService ScheduledService
}

type ScheduledService interface {
	Run()
}

func (service SchedulerServiceImpl) Start() {
	for {
		service.ScheduledService.Run()
		service.waitTillNextJobRun()
	}
}

func (service SchedulerServiceImpl) waitTillNextJobRun() {
	jobIntervalMinutes := service.IntervalMinutes
	fmt.Println("Waiting", jobIntervalMinutes, "minutes till the next job run")
	jobInterval := time.Minute * time.Duration(jobIntervalMinutes)
	time.Sleep(jobInterval)
}
