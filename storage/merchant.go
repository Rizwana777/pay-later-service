package storage

import (
	"context"
	"fmt"
	"pay-later-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveMerchant inserts a new document for mercahnt in database.
func (r *Repo) SaveMerchant(m *model.Merchant) error {
	collection := r.client.Database(dbName).Collection(merchantCollection)
	_, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// GetMerchant retrieves an existing document for merchant based on email from database.
func (r *Repo) GetMerchant(name string) (*model.Merchant, error) {
	var user model.Merchant
	collection := r.client.Database(dbName).Collection(merchantCollection)
	filter := primitive.M{"name": name}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateMerchant updates the merchant data in database based on name.
func (s *Repo) UpdateMerchant(m model.Merchant) error {
	collection := s.client.Database(dbName).Collection(merchantCollection)
	filter := bson.M{"name": m.Name}
	f := bson.D{}
	f = append(f, bson.E{"totalDiscount", m.TotalDiscount})
	f = append(f, bson.E{"discount", m.Discount})
	merchant, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": f})
	if err != nil {
		return err
	}

	if merchant.MatchedCount != 1 {
		return err
	}

	return nil
}
