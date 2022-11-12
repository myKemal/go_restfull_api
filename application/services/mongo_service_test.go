package services

import (
	"errors"
	"testing"
	"time"

	mockrepo "github.com/myKemal/go_restfull_api/application/mocks"

	"github.com/myKemal/go_restfull_api/application/common"
	"github.com/myKemal/go_restfull_api/application/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoService_GetRecords(t *testing.T) {
	//given

	startDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-01")
	endDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-02")

	repo := new(mockrepo.MongoRepository)
	recordsService := NewMongoServicewtihRepo(repo)

	repo.On("GetRecordsWith", 100, 200, startDate, endDate).
		Return([]models.MongoRecord{{
			Id:         primitive.ObjectID{100},
			Key:        "200",
			Value:      "300",
			CreatedAt:  startDate,
			TotalCount: 400,
		}}, nil)

	//when
	records, getRecordsErr := recordsService.GetRecords(100, 200, startDate, endDate)

	//then
	assert.NoError(t, getRecordsErr)
	assert.Equal(t, len(records), 1)
	assert.Equal(t, records[0].Key, "200")
	assert.Equal(t, records[0].CreatedAt, startDate)
	assert.Equal(t, records[0].TotalCount, 400)
}

func TestMongoRecordsService_RepositoryError(t *testing.T) {
	//given
	repo := new(mockrepo.MongoRepository)
	recordsService := NewMongoServicewtihRepo(repo)
	startDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-01")
	endDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-02")
	unexpectedError := errors.New("connection timeout")

	repo.On("GetRecordsWith", 100, 200, startDate, endDate).
		Return(nil, unexpectedError)

	//when
	records, getRecordsErr := recordsService.GetRecords(100, 200, startDate, endDate)

	//then
	assert.Empty(t, records)
	assert.Error(t, getRecordsErr)
	assert.Equal(t, getRecordsErr, unexpectedError)
}
