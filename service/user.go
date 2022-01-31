package service

import (
	"fmt"
	"pay-later-service/model"
	"pay-later-service/storage"
	"strconv"
	"strings"
)

func (c *Context) CreateUser(user string) {
	var users model.User

	u := strings.Split(user, " ")

	users.Name = u[2]
	users.Email = u[3]
	creditLimit, err := strconv.ParseFloat(u[4], 64)
	if err != nil {
		fmt.Println(err)
	}
	users.CreditLimit = creditLimit
	users.Balance = users.CreditLimit

	if _, ok := validMailAddress(users.Email); !ok {
		fmt.Println("invalid email")
		return
	}

	repo := storage.NewRepository(c.DBClient)

	_, err = repo.GetUser(users.Name)
	if err == nil {
		fmt.Println("user with name " + users.Name + " already exists")
		return
	}

	err = repo.SaveUser(&users)
	if err != nil {
		fmt.Println("unable to save user", err)
	}

	fmt.Printf("%s(%0.2f)\n", users.Name, users.CreditLimit)

}
