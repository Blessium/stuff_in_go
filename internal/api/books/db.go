package books

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookDB struct {
	id        primitive.ObjectID `bson:"_id"`
	ISBN      string             `bson:"ISBN"`
	Title     string             `bson:"Title"`
	Author    string             `bson:"Author"`
	Published primitive.DateTime `bson:"Published"`
	Pages     uint               `bson:"Pages"`
}

func (b *BookDB) SetPublished(p time.Time) {
	b.Published = primitive.NewDateTimeFromTime(p)
}

type IRepository interface {
	Create(ctx context.Context, book BookDB) (BookDB, error)
	Update(ctx context.Context, book BookDB) (BookDB, error)
	Delete(ctx context.Context, isbn string) error
	Get(ctx context.Context, isbn string) (BookDB, error)
}

type MongoRepository struct {
	db *mongo.Collection
}

func GetMongoRepository(db *mongo.Collection) MongoRepository {
	return MongoRepository{
		db: db,
	}
}

func (b MongoRepository) Create(ctx context.Context, book BookDB) (BookDB, error) {
	_, err := b.db.InsertOne(ctx, book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b MongoRepository) Update(ctx context.Context, book BookDB) (BookDB, error) {
	filter := bson.M{"_id": book.ISBN}
	update := bson.M{"$set": bson.M{
		"Title":     book.Title,
		"Published": book.Published,
		"Pages":     book.Pages,
		"Author":    book.Author,
	}}
	res, err := b.db.UpdateOne(ctx, filter, update, nil)
	if err != nil {
		return book, err
	}

	if res.MatchedCount == 0 {
		return book, errors.New("element doesn't exist")
	}
	return book, nil
}

func (b MongoRepository) Delete(ctx context.Context, isbn string) error {
	return nil
}

func (b MongoRepository) Get(ctx context.Context, isbn string) (BookDB, error) {
	book := BookDB{}
	return book, nil
}
