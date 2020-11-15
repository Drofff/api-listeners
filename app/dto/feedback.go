package dto

import (
	"api-listeners/app/util"
)

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

type FeedbackDto struct {
	ID int64 `json:"id"`
	Rate string `json:"customerFeedback"`
	Comment string `json:"ratingComment"`
	Name string `json:"customerName"`
	PhoneNumber string `json:"selfServiceUserMobile"`
}

func (dto FeedbackDto) IsNotComplete() bool {
	return util.IsBlank(dto.Rate) || util.IsBlank(dto.Name) || util.IsBlank(dto.PhoneNumber)
}

const zeroValueStr = "<unknown>"

func (dto FeedbackDto) String() string {
	r, err := util.ToJson(dto)
	if err != nil {
		return zeroValueStr
	}
	s, err := util.ReadAsString(r)
	if err != nil {
		return zeroValueStr
	}
	return s
}