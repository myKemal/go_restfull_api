package services

import (
	"time"

	db "github.com/myKemal/go_restfull_api/application/dbClient"
	"github.com/myKemal/go_restfull_api/application/dto"

	repo "github.com/myKemal/go_restfull_api/application/repositories"
)

type MongoService interface {
	GetRecords(minCount, maxCount int, startDate, endDate time.Time) ([]dto.MongoRecordResponse, error)
}

type mongoService struct {
	recordsRepository repo.MongoRepository
}

func NewMongoServicewtihRepo(repo repo.MongoRepository) MongoService {
	return &mongoService{
		recordsRepository: repo,
	}
}

func NewMongoService(client db.MongoClient) MongoService {
	return &mongoService{
		recordsRepository: repo.NewMongoRepository(client),
	}
}

func (m *mongoService) GetRecords(minCount, maxCount int, startDate, endDate time.Time) ([]dto.MongoRecordResponse, error) {
	records, recordsErr := m.recordsRepository.GetRecordsWith(minCount, maxCount, startDate, endDate)
	if recordsErr != nil {
		return nil, recordsErr
	}
	recordResponse := []dto.MongoRecordResponse{}
	for _, v := range records {
		recordResponse = append(recordResponse, dto.MongoRecordResponse{Key: v.Key, CreatedAt: v.CreatedAt, TotalCount: v.TotalCount})
	}
	return recordResponse, nil
}
