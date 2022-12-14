package dto

import "time"

type MongoRecordsResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"msg"`
	Records []MongoRecordResponse `json:"records"`
}

type MongoRecordResponse struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type MongoGetRecordsRequest struct {
	StartDate string `json:"startDate" example:"2016-10-02" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate" example:"2021-01-02" validate:"required,datetime=2006-01-02,gtefield=EndDate"`
	MinCount  int    `json:"minCount" example:"2800" validate:"required,gte=0"`
	MaxCount  int    `json:"maxCount" example:"3000" validate:"required,gte=0,gtefield=MinCount"`
}
