package storage

import (
	"pay-later-service/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName             = "pay_later"
	userCollection     = "users"
	merchantCollection = "merchant"
	transaction        = "transaction"
)

// Repo implements the Repository interface.
type Repo struct {
	client *mongo.Client
}

var _ Repository = &Repo{}

// NewRepository returns Repo object.
func NewRepository(client *mongo.Client) *Repo {
	return &Repo{client}
}

type Repository interface {
	SaveUser(u *model.User) error
	GetUser(email string) (*model.User, error)
	UpdateUser(fg model.User) error
	ListUsers() ([]model.User, error)
	GetAllUsers() ([]model.User, error)
	SaveMerchant(m *model.Merchant) error
	GetMerchant(email string) (*model.Merchant, error)
	UpdateMerchant(fg model.Merchant) error
}
