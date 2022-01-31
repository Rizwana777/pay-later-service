package storage

import (
	"context"
	"fmt"
	"pay-later-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveUser inserts a new document for user in database.
func (r *Repo) SaveUser(u *model.User) error {
	u.ID = primitive.NewObjectID().Hex()
	collection := r.client.Database(dbName).Collection(userCollection)
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// GetUser retrieves an existing document for user based on email from database.
func (r *Repo) GetUser(name string) (*model.User, error) {
	var u model.User
	collection := r.client.Database(dbName).Collection(userCollection)
	filter := primitive.M{"name": name}
	err := collection.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// ListUsers fetches user(s) from database
func (r *Repo) ListUsers() ([]model.User, error) {
	var users []model.User
	findOptions := options.Find()
	collection := r.client.Database(dbName).Collection(userCollection)

	filter := primitive.M{}
	filter["balance"] = 0

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return users, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		user := model.User{}
		err = cur.Decode(&user)

		if err == nil {
			users = append(users, user)
		} else {
			return users, err
		}
	}

	return users, nil
}

// GetAllUsers fetches user(s) from database
func (r *Repo) GetAllUsers() ([]model.User, error) {
	var users []model.User

	findOptions := options.Find()

	collection := r.client.Database(dbName).Collection(userCollection)
	filter := primitive.M{}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return users, err
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		user := model.User{}
		err = cur.Decode(&user)
		if err == nil {
			users = append(users, user)
		} else {
			return users, err
		}
	}
	return users, nil
}

// UpdateUser updates the user database based on name.
func (s *Repo) UpdateUser(u model.User) error {
	collection := s.client.Database(dbName).Collection(userCollection)
	filter := bson.M{"name": u.Name}
	f := bson.D{}
	f = append(f, bson.E{"due", u.Due})
	f = append(f, bson.E{"balance", u.Balance})

	user, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": f})
	if err != nil {
		return err
	}

	if user.MatchedCount != 1 {
		return err
	}

	return nil
}
