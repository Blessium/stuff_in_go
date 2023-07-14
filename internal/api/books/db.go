package books

import (
    "context"
    "github.com/blessium/metricsgo/internal/db"
    "go.mongodb.org/mongo-driver/mongo"
)

type BookDB struct {

}

type IRepository interface {
	Create(ctx context.Context, book BookDB) (BookDB, error)
	Update(ctx context.Context, book BookDB) (BookDB, error)
	Delete(ctx context.Context, isbn string) error
	Get(ctx context.Context, isbn string) (BookDB, error)
}

type MongoRepository struct {
    db *mongo.Client
} 

func GetMongoRepository() (IRepository, error) {
    db, err := db.GetMongoDB()
    if err != nil {
        return nil, err
    }

    return MongoRepository {
        db: db,
    }, nil
}

func (b MongoRepository) Create(ctx context.Context, book BookDB) (BookDB, error) {
    return book, nil
}

func (b MongoRepository) Update(ctx context.Context, book BookDB) (BookDB, error) {
    return book, nil
}

func (b MongoRepository) Delete(ctx context.Context, isbn string) error {
    return nil
}

func (b MongoRepository) Get(ctx context.Context, isbn string) (BookDB, error) {
    book := BookDB {}
    return book, nil
}
