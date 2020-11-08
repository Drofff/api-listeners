package service

import (
	"api-listeners/app/dto"
	"fmt"
)

type FeedbacksProcessor interface {
	ProcessFeedbacks(feedbacks *dto.FeedbacksDto)
}

type LoggingFeedbacksProcessor struct {

}

func (service *LoggingFeedbacksProcessor) ProcessFeedbacks(feedbacks *dto.FeedbacksDto) {
	fmt.Println(feedbacks.Response)
}