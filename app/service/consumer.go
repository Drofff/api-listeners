package service

import (
	"api-listeners/app/dto"
	"api-listeners/app/util"
	"fmt"
	"log"
	"time"
)

type FeedbacksConsumer struct {
	GetFeedbacksUrl string
	RefreshToken string
	FeedbacksPageSize int64
	NextPageDelayMinutes int64
	AuthorizationService AuthorizationService
	StateHolder StateHolder
	FeedbacksProcessor FeedbacksProcessor
}

const lastFetchTimeState = "last-fetch"
const lastFetchedPageState = "last-page"

func (consumer *FeedbacksConsumer) Run() {
	lft, err := consumer.getLastFetchTime()
	if err != nil {
		return
	}
	defer consumer.updateLastFetchTime()
	defer consumer.clearLastFetchedPage()
	for {
		lfp := consumer.getLastFetchedPage()
		nfp := lfp + 1
		feedbacksUrl := consumer.getFeedbacksUrl(nfp, lft)
		feedbacks := &dto.FeedbacksDto{}
		err = util.DoGetWithToken(feedbacksUrl, consumer.authorizationToken(), feedbacks)
		if err != nil {
			fmt.Println("Couldn't fetch feedbacks because of: ", err)
			return
		}
		consumer.FeedbacksProcessor.ProcessFeedbacks(feedbacks)
		if !feedbacks.HasNextPage() {
			break
		}
		consumer.updateLastFetchedPage(nfp)
		consumer.waitToConsumeNextPage()
	}
}

func (consumer *FeedbacksConsumer) getLastFetchTime() (int64, error) {
	lastFetch, err := consumer.StateHolder.GetIntState(lastFetchTimeState)
	if err != nil {
		switch err.(type) {
			case UnknownPrefixError:
				fmt.Println("WARN: No last fetch time has been set. Setting the current time and skipping feedbacks fetch")
				consumer.updateLastFetchTime()
				return -1, err
			default:
				log.Fatal(err)
		}
	}
	return lastFetch, nil
}

func (consumer *FeedbacksConsumer) updateLastFetchTime() {
	currTime := time.Now().Unix()
	consumer.StateHolder.SetIntState(lastFetchTimeState, currTime)
}

func (consumer *FeedbacksConsumer) clearLastFetchedPage() {
	consumer.StateHolder.SetIntState(lastFetchedPageState, -1)
}

func (consumer *FeedbacksConsumer) getLastFetchedPage() int64 {
	lastPage, err := consumer.StateHolder.GetIntState(lastFetchedPageState)
	if err != nil {
		switch err.(type) {
		case UnknownPrefixError:
			return -1
		default:
			log.Fatal(err)
		}
	}
	return lastPage
}

func (consumer *FeedbacksConsumer) getFeedbacksUrl(page int64, from int64) string {
	to := time.Now().Unix()
	return fmt.Sprintf("%v?page=%v&size=%v&startDate=%v&endDate%v", consumer.GetFeedbacksUrl,
		page, consumer.FeedbacksPageSize, from, to)
}

func (consumer *FeedbacksConsumer) authorizationToken() string {
	authzToken, err := consumer.AuthorizationService.GetAuthorizationToken()
	if err != nil {
		log.Fatal("Error while getting an authorization token: ", err)
	}
	return authzToken
}

func (consumer *FeedbacksConsumer) updateLastFetchedPage(page int64) {
	consumer.StateHolder.SetIntState(lastFetchedPageState, page)
}

func (consumer *FeedbacksConsumer) waitToConsumeNextPage() {
	nextPageDelay := consumer.NextPageDelayMinutes
	fmt.Printf("Waiting %v minutes to consume next feedbacks page\n", nextPageDelay)
	nextPageDelayTime := time.Duration(nextPageDelay) * time.Minute
	time.Sleep(nextPageDelayTime)
}
