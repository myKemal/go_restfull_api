package repositories

import (
	"context"
	"time"

	db "github.com/myKemal/go_restfull_api/application/dbClient"

	"github.com/myKemal/go_restfull_api/application/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	GetRecordsWith(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]models.MongoRecord, error)
}

type mongoRepository struct {
	mongoClient db.MongoClient
}

func NewMongoRepository(client db.MongoClient) MongoRepository {
	return &mongoRepository{
		mongoClient: client,
	}
}

// GetRecordsWith Returns Records from mongodb with the filters provided as minCount, maxCount, startDate and endDate or
// returns an error
func (r *mongoRepository) GetRecordsWith(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]models.MongoRecord, error) {

	var query primitive.M
	var queryAndArray primitive.A

	queryAndArray = append(queryAndArray, bson.M{"createdAt": bson.M{"$gte": startDate.UTC().Format("2006-01-02T15:04:05-0700")}})
	queryAndArray = append(queryAndArray, bson.M{"createdAt": bson.M{"$lte": endDate.UTC().Format("2006-01-02T15:04:05-0700")}})
	queryAndArray = append(queryAndArray, bson.M{"totalCount": bson.M{"$gte": minCount}})
	queryAndArray = append(queryAndArray, bson.M{"totalCount": bson.M{"$lte": maxCount}})

	query = bson.M{"$and": queryAndArray}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	cursor, findErr := r.getCollection().Find(ctx, query)

	if findErr != nil {
		defer cancel()
		return nil, findErr
	}

	defer func(_cursor *mongo.Cursor) {
		cancel()
		_ = _cursor.Close(ctx)
	}(cursor)

	var mongoRecords []models.MongoRecord
	findAllErr := cursor.All(context.Background(), &mongoRecords)
	if findAllErr != nil {
		return nil, findAllErr
	}
	return mongoRecords, nil
}

// getCollection Returns a records collection from the database
func (r *mongoRepository) getCollection() *mongo.Collection {
	return r.mongoClient.GetCollection()
}
