package services

import (
	"context"
	"errors"
	"time"

	"fatawa-api/pkg/models"
	"fatawa-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FatwaServiceImpl struct {
	fatwaCollection *mongo.Collection
	ctx             context.Context
}

func NewFatwaService(fatwaCollection *mongo.Collection, ctx context.Context) FatwaService {
	return &FatwaServiceImpl{fatwaCollection, ctx}
}

func (p *FatwaServiceImpl) CreateFatwa(fatwa *models.Fatwa) (*models.FatwaDb, error) {
	fatwa.CreatedAt = time.Now()
	fatwa.UpdatedAt = fatwa.CreatedAt
	res, err := p.fatwaCollection.InsertOne(p.ctx, fatwa)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("post with that title already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := p.fatwaCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}

	var newFatwa *models.FatwaDb
	query := bson.M{"_id": res.InsertedID}
	if err = p.fatwaCollection.FindOne(p.ctx, query).Decode(&newFatwa); err != nil {
		return nil, err
	}

	return newFatwa, nil
}

func (p *FatwaServiceImpl) UpdateFatwa(id string, data *models.FatwaDb) (*models.FatwaDb, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.fatwaCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedFatwa *models.FatwaDb

	if err := res.Decode(&updatedFatwa); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedFatwa, nil
}

func (p *FatwaServiceImpl) FindFatwaById(id string) (*models.FatwaDb, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var fatwa *models.FatwaDb

	if err := p.fatwaCollection.FindOne(p.ctx, query).Decode(&fatwa); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return fatwa, nil
}

func (p *FatwaServiceImpl) FindFatawa(page int, limit int) ([]*models.FatwaDb, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}

	cursor, err := p.fatwaCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, p.ctx)

	var fatawa []*models.FatwaDb

	for cursor.Next(p.ctx) {
		fatwa := &models.FatwaDb{}
		err := cursor.Decode(fatwa)

		if err != nil {
			return nil, err
		}

		fatawa = append(fatawa, fatwa)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(fatawa) == 0 {
		return []*models.FatwaDb{}, nil
	}

	return fatawa, nil
}

func (p *FatwaServiceImpl) DeleteFatwa(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.fatwaCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
