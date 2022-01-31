package model

type User struct {
	ID          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Email       string  `bson:"email"`
	CreditLimit float64 `bson:"creditLimit"`
	Due         float64 `bson:"due,omitempty"`
	Balance     float64 `bson:"balance"`
}
