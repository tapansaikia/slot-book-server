package storage

type MongoStorage struct{}

func NewMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

func (ms *MongoStorage) Get(id int) {
	// fetch from mongodb
}
