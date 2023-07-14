package db

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
)

const uri=""

func GetMongoDB() (*mongo.Client, error) {

    serverApi := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        return nil, err
    }

    return client, nil
}  


