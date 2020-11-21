package main

import (
	"api-listeners/app/service"
	"fmt"
)

func main() {
	fmt.Println("Starting API Listeners")
	configService := &service.EmbeddedFileConfigService{ConfigFilePath: "application.properties"}
	feedbacksConsumer := &service.FeedbacksConsumer{
		GetFeedbacksUrl: configService.GetProp("orty-api.get-feedbacks.url"),
		RefreshToken: configService.GetProp("orty-api.refresh-token"),
		FeedbacksPageSize: configService.GetIntProp("feedbacks.page-size"),
		NextPageDelayMinutes: configService.GetIntProp("feedbacks.next-page-delay-minutes"),
		AuthorizationService: newAuthorizationService(configService),
		StateHolder: &service.FileStateHolder{FilePath: configService.GetProp("state-holder.file-path")},
		FeedbacksProcessor: newFeedbacksProcessor(configService),
	}
	scheduler := service.SchedulerServiceImpl{
		IntervalMinutes: configService.GetIntProp("job-run.interval.minutes"),
		ScheduledService: feedbacksConsumer,
	}
	scheduler.Start()
}

func newAuthorizationService(configService service.ConfigService) service.AuthorizationService {
	return &service.JwtAuthorizationService{
		LoginUrl: configService.GetProp("orty-api.login.url"),
		Username: configService.GetProp("orty-api.username"),
		Password: configService.GetProp("orty-api.password"),
		TokenTimeToLiveMinutes: configService.GetIntProp("orty-api.jwt.time-to-live-minutes"),
	}
}

func newFeedbacksProcessor(configService service.ConfigService) service.FeedbacksProcessor {
	return &service.BotApiFeedbacksProcessor{
		SendFeedbacksUrl: configService.GetProp("bot-api.send-feedbacks.url"),
	}
}