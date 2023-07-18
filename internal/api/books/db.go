package books

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	GetAll(ctx context.Context) ([]BookDB, error)
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
	filter := bson.M{"_id": book.id}
	update := bson.M{"$set": bson.M{
		"ISBN":      book.ISBN,
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
	filter := bson.M{"isbn": isbn}
	result, err := b.db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("Element doesn't exist")
	}
	return nil
}

func (b MongoRepository) Get(ctx context.Context, isbn string) (BookDB, error) {
	book := BookDB{}
	filter := bson.M{"isbn": isbn}
	err := b.db.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (b MongoRepository) GetAll(ctx context.Context) ([]BookDB, error) {
    findOptions := options.Find()
	cursor, err := b.db.Find(ctx, bson.M{}, findOptions)
	if err != nil {
        return []BookDB{}, err
    }

    var books []BookDB
    if err := cursor.All(ctx, &books); err != nil {
        return []BookDB{}, err
    }

    return books, nil
}
