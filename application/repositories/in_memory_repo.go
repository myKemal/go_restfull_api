package repositories

import (
	"errors"

	db "github.com/myKemal/go_restfull_api/application/dbClient"
)

const KeyAlreadyExistError = "Key already exist."
const KeyNotExistError = "Key not exist."

type InMemoryRecordsRepository interface {
	Get(key string) (string, error)
	Create(key, value string) error
}

type inMemoryRecordsRepository struct {
	client db.InMemoryClient
}

func NewInMemoryRecordsRepository(client db.InMemoryClient) InMemoryRecordsRepository {
	return &inMemoryRecordsRepository{
		client: client,
	}
}

// Get Returns value of a provided key from database or throws an error if not exist.
func (r *inMemoryRecordsRepository) Get(key string) (string, error) {
	record := r.client.GetSync(key)
	if len(record) == 0 {
		return "", errors.New(KeyNotExistError)
	}
	return record, nil
}

// Create Returns creates a key-value pair in the database or throws an error if already exist.
func (r *inMemoryRecordsRepository) Create(key, value string) error {
	record := r.client.GetSync(key)
	if len(record) != 0 {
		return errors.New(KeyAlreadyExistError)
	}
	r.client.CreateSync(key, value)
	return nil
}
