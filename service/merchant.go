package service

import (
	"fmt"
	"pay-later-service/model"
	"pay-later-service/storage"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Context) CreateMerchant(mr string) {
	var merchant model.Merchant

	m := strings.Split(mr, " ")

	merchant.ID = primitive.NewObjectID().Hex()
	merchant.Name = m[2]
	merchant.Email = m[3]
	discount, err := strconv.ParseFloat(m[4], 64)
	if err != nil {
		fmt.Println(err)
	}
	merchant.Discount = discount

	if _, ok := validMailAddress(merchant.Email); !ok {
		fmt.Println("invalid email")
		return
	}

	repo := storage.NewRepository(c.DBClient)

	_, err = repo.GetMerchant(merchant.Name)
	if err == nil {
		fmt.Println("merchant with name " + merchant.Name + " already exists")
		return
	}

	err = repo.SaveMerchant(&merchant)
	if err != nil {
		fmt.Println("unable to save merchant", err)
		return
	}

	fmt.Printf("%s(%0.2f %%)\n", merchant.Name, merchant.Discount)

}

func (c *Context) UpdateMerchant(mr string) {
	m := strings.Split(mr, " ")
	repo := storage.NewRepository(c.DBClient)
	merchant, err := repo.GetMerchant(m[2])
	if err != nil {
		fmt.Println("unable to get merchant from database")
		return
	}
	discount, err := strconv.ParseFloat(m[3], 64)
	if err != nil {
		fmt.Println(err)
	}
	merchant.Discount = discount

	err = repo.UpdateMerchant(*merchant)
	if err != nil {
		fmt.Println("error updating merchant", err)
		return
	}
	fmt.Println("updated!")

}
