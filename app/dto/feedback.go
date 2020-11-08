package dto

type FeedbacksDto struct {
	Response struct {
		Data struct {
			Last bool `json:"last"`
			FeedbackList []FeedbackDto `json:"items"`
		} `json:"data"`
	} `json:"response"`
}

func (dto *FeedbacksDto) HasNextPage() bool {
	return !dto.Response.Data.Last
}

type FeedbackDto map[string]interface{}
