package services

import (
	"net/http"
	"strings"

	"github.com/myKemal/go_restfull_api/application/common"
	db "github.com/myKemal/go_restfull_api/application/dbClient"
	inmemory "github.com/myKemal/go_restfull_api/application/repositories"
)

type InMemoryRecordsService interface {
	Create(key string, value string) error
	Get(key string) (string, error)
}

type inMemoryRecordsService struct {
	recordsRepository inmemory.InMemoryRecordsRepository
}

func NewInMemoryRecordsService(client db.InMemoryClient) InMemoryRecordsService {
	return &inMemoryRecordsService{
		recordsRepository: inmemory.NewInMemoryRecordsRepository(client),
	}
}

func NewInMemoryRecordsServiceWith(repository inmemory.InMemoryRecordsRepository) InMemoryRecordsService {
	return &inMemoryRecordsService{
		recordsRepository: repository,
	}
}

func (i *inMemoryRecordsService) Create(key string, value string) error {
	createErr := i.recordsRepository.Create(key, value)
	if createErr != nil {
		if strings.EqualFold(createErr.Error(), inmemory.KeyAlreadyExistError) {
			return common.NewApiError(http.StatusConflict, 1, "Record with key already exist.")
		}
		return createErr
	}
	return nil
}

func (i *inMemoryRecordsService) Get(key string) (string, error) {
	record, getErr := i.recordsRepository.Get(key)
	if getErr != nil {
		if strings.EqualFold(getErr.Error(), inmemory.KeyNotExistError) {
			return "", common.NewApiError(http.StatusOK, 1, "Record not found")
		}
		return "", getErr
	}
	return record, nil
}
