package service

import (
	"fmt"
	"pay-later-service/model"
	"pay-later-service/storage"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Context) CreateTransaction(tr string) {
	var t model.Transaction

	m := strings.Split(tr, " ")

	t.ID = primitive.NewObjectID().Hex()
	t.UserName = m[2]
	t.MerchantName = m[3]
	amount, err := strconv.ParseFloat(m[4], 64)
	if err != nil {
		fmt.Println(err)
	}
	t.Amount = amount

	repo := storage.NewRepository(c.DBClient)
	user, err := repo.GetUser(t.UserName)
	if err != nil {
		fmt.Println("unable to get user from database")
		return
	}

	if user.Balance <= 0 || t.Amount > user.Balance {
		fmt.Println("credit limit reached")
		return
	}

	balance := user.Balance - t.Amount
	user.Due = user.CreditLimit - balance
	user.Balance = balance

	err = repo.UpdateUser(*user)
	if err != nil {
		fmt.Println("error updating user", err)
		return
	}

	merchant, err := repo.GetMerchant(t.MerchantName)
	if err != nil {
		fmt.Println("unable to get merchant from database")
		return
	}

	// total discount a merchant given to a user
	totalDiscount := amount * float64(merchant.Discount) / 100
	merchant.TotalDiscount = merchant.TotalDiscount + totalDiscount

	err = repo.UpdateMerchant(*merchant)
	if err != nil {
		fmt.Println("error updating merchant", err)
		return
	}

	fmt.Println("success!")

}
