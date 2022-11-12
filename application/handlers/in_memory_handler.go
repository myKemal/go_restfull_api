package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/myKemal/go_restfull_api/application/common"
	db "github.com/myKemal/go_restfull_api/application/dbClient"
	"github.com/myKemal/go_restfull_api/application/models"
	"github.com/myKemal/go_restfull_api/application/server"
	"github.com/myKemal/go_restfull_api/application/services"
)

type InMemoryHandler interface {
	Create(request *server.Request, response *server.Response)
	GetRecords(request *server.Request, response *server.Response)
}

type inMemoryHandler struct {
	recordsService services.InMemoryRecordsService
}

func NewInMemoryHandler(client db.InMemoryClient) InMemoryHandler {
	return &inMemoryHandler{
		recordsService: services.NewInMemoryRecordsService(client),
	}
}

func (i inMemoryHandler) Create(req *server.Request, res *server.Response) {
	var requestBody models.InMemoryCreateRecordRequest
	unmarshallErr := json.Unmarshal(req.Body, &requestBody)
	if unmarshallErr != nil {
		res.Error = common.NewBadRequestError()
		return
	}
	validateErr := common.Validate.Struct(requestBody)
	if validateErr != nil {
		res.Error = common.NewBadRequestErrorWithMessage(common.TranslateValidationErrors(validateErr))
		return
	}
	createErr := i.recordsService.Create(requestBody.Key, requestBody.Value)
	if createErr != nil {
		res.Error = createErr
		return
	}
	res.StatusCode = http.StatusCreated
	res.Body = models.InMemoryRecordResponse{Key: requestBody.Key, Value: requestBody.Value}
}

func (i inMemoryHandler) GetRecords(req *server.Request, res *server.Response) {
	var requestData models.InMemoryGetRecordRequest
	value, paramExist := req.Parameters["key"]
	if !paramExist {
		res.Error = common.NewBadRequestError()
		return
	}
	requestData.Key = value[0]
	validateErr := common.Validate.Struct(requestData)
	if validateErr != nil {
		res.Error = common.NewBadRequestErrorWithMessage(common.TranslateValidationErrors(validateErr))
		return
	}

	record, getErr := i.recordsService.Get(requestData.Key)
	if getErr != nil {
		res.Error = getErr
		return
	}
	res.Body = models.InMemoryRecordResponse{Key: value[0], Value: record}
}
