package model

type Transaction struct {
	ID           string  `bson:"_id"`
	UserName     string  `bson:"userName"`
	MerchantName string  `bson:"merchantName"`
	Amount       float64 `bson:"amount"`
}
