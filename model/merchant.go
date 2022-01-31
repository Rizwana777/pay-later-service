package model

type Merchant struct {
	ID            string  `bson:"_id"`
	Name          string  `bson:"name"`
	Email         string  `bson:"email"`
	Discount      float64 `bson:"discount"`
	TotalDiscount float64 `bson:"totalDiscount,omitempty"`
}
