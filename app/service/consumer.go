package service

import (
	"api-listeners/app/cache"
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
	timeRun := util.UnixMillis()
	lft, err := consumer.getLastFetchTime()
	if err != nil {
		return
	}
	consumer.updateLastFetchTime()
	defer consumer.clearLastFetchedPage()
	for {
		feedbacks := &dto.FeedbacksDto{}
		err = consumer.nextFeedbacksPageFromTo(lft, timeRun, feedbacks)
		if err != nil {
			fmt.Println("Couldn't fetch feedbacks because of: ", err)
			return
		}
		consumer.FeedbacksProcessor.ProcessFeedbacks(feedbacks)
		if !feedbacks.HasNextPage() {
			break
		}
		consumer.incrementLastFetchedPage()
		consumer.waitToConsumeNextPage()
	}
	cache.Submit()
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
	currTime := util.UnixMillis()
	consumer.StateHolder.SetState(lastFetchTimeState, currTime)
}

func (consumer *FeedbacksConsumer) clearLastFetchedPage() {
	consumer.StateHolder.SetState(lastFetchedPageState, -1)
}

func (consumer *FeedbacksConsumer) nextFeedbacksPageFromTo(from, to int64, res *dto.FeedbacksDto) error {
	page := consumer.getNextPageToFetch()
	feedbacksUrl := consumer.getFeedbacksUrl(page, from, to)
	return util.DoGetWithToken(feedbacksUrl, consumer.authorizationToken(), res)
}

func (consumer *FeedbacksConsumer) getFeedbacksUrl(page, from, to int64) string {
	return fmt.Sprintf("%v?page=%v&size=%v&startDate=%v&endDate=%v", consumer.GetFeedbacksUrl,
		page, consumer.FeedbacksPageSize, from, to)
}

func (consumer *FeedbacksConsumer) authorizationToken() string {
	authzToken, err := consumer.AuthorizationService.GetAuthorizationToken()
	if err != nil {
		log.Fatal("Error while getting an authorization token: ", err)
	}
	return authzToken
}

func (consumer *FeedbacksConsumer) incrementLastFetchedPage() {
	nextPage := consumer.getNextPageToFetch()
	consumer.StateHolder.SetState(lastFetchedPageState, nextPage)
}

func (consumer *FeedbacksConsumer) getNextPageToFetch() int64 {
	return consumer.getLastFetchedPage() + 1
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

func (consumer *FeedbacksConsumer) waitToConsumeNextPage() {
	nextPageDelay := consumer.NextPageDelayMinutes
	fmt.Printf("Waiting %v minutes to consume next feedbacks page\n", nextPageDelay)
	nextPageDelayTime := time.Duration(nextPageDelay) * time.Minute
	time.Sleep(nextPageDelayTime)
}
