// myapp/internal/db/mongodb/MongoDBFactory.go
package mongodb

func GetNewMongoDBManager() (*MongoDBManager, error) {
	client, err := NewMongoDBClient()
	if err != nil {
		return nil, err
	}

	manager := NewMongoDBManager(client)
	return manager, nil
}
