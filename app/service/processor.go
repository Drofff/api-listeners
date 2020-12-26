package service

import (
	"api-listeners/app/cache"
	"api-listeners/app/dto"
	"api-listeners/app/util"
	"fmt"
	"io"
)

type FeedbacksProcessor interface {
	ProcessFeedbacks(feedbacks *dto.FeedbacksDto)
}

type BotApiFeedbacksProcessor struct {
	SendFeedbacksUrl string
}

func (service *BotApiFeedbacksProcessor) ProcessFeedbacks(feedbacks *dto.FeedbacksDto) {
	feedbackList := feedbacks.Response.Data.FeedbackList
	if len(feedbackList) == 0 {
		fmt.Println("INFO: 0 feedbacks received")
		return
	}
	for _, feedback := range feedbackList {
		service.processFeedback(feedback)
	}
}

func (service *BotApiFeedbacksProcessor) processFeedback(feedback dto.FeedbackDto) {
	if feedback.IsNotComplete() {
		fmt.Printf("INFO: skipping incomplete feedback %v\n", feedback)
		return
	}
	if cache.IsDuplicatedID(feedback.ID) {
		fmt.Printf("INFO: skipping duplicated feedback %v with id '%v'\n", feedback, feedback.ID)
		return
	}
	botApiFeedbackDto := asBotApiFeedbackDto(feedback)
	err := util.DoPostJson(service.SendFeedbacksUrl, botApiFeedbackDto, &struct{}{})
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR: can not send feedback %v because of - %v\n", feedback, err)
	}
	cache.SaveID(feedback.ID)
}

func asBotApiFeedbackDto(feedback dto.FeedbackDto) interface{} {
	return struct {
		Rate string `json:"rate"`
		Comment string `json:"comment"`
		Name string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
	}{
		Rate: feedback.Rate,
		Comment: feedback.Comment,
		Name: feedback.Name,
		PhoneNumber: feedback.PhoneNumber,
	}
}
