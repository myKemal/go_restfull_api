package models

type InMemoryCreateRecordRequest struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type InMemoryGetRecordRequest struct {
	Key string `json:"key" validate:"required"`
}

type InMemoryRecordResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
