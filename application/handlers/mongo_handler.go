package handlers

import (
	"encoding/json"
	"time"

	"github.com/myKemal/go_restfull_api/application/common"
	db "github.com/myKemal/go_restfull_api/application/dbClient"
	"github.com/myKemal/go_restfull_api/application/dto"
	"github.com/myKemal/go_restfull_api/application/server"
	"github.com/myKemal/go_restfull_api/application/services"
)

type MongoHandler interface {
	GetRecords(request *server.Request, response *server.Response)
}

type mongoHandler struct {
	recordsService services.MongoService
}

func NewMongoHandler(client db.MongoClient) MongoHandler {
	return &mongoHandler{
		recordsService: services.NewMongoService(client),
	}
}

// GetRecords
// @Summary fetch data from mongodb
// @Tags mongo
// Description fetches the data in the provided MongoDB collection and returns the results in the requested format.
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.MongoRecordsResponse
// @Failure 400 {object} common.ApiError
// @Router /api/v1/mongo [post]
// @Param Request body dto.MongoGetRecordsRequest true "Filter for the request"
func (m mongoHandler) GetRecords(req *server.Request, res *server.Response) {
	var requestBody dto.MongoGetRecordsRequest
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
	startDate, _ := time.Parse(common.DefaultTimeFormat, requestBody.StartDate)
	endDate, _ := time.Parse(common.DefaultTimeFormat, requestBody.EndDate)
	records, recordsErr := m.recordsService.GetRecords(requestBody.MinCount, requestBody.MaxCount, startDate, endDate)
	if recordsErr != nil {
		res.Error = recordsErr
		return
	}
	res.Body = dto.MongoRecordsResponse{Code: 0, Message: "Success", Records: records}
}
