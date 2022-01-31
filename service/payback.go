package service

import (
	"fmt"
	"pay-later-service/storage"
	"strconv"
	"strings"
)

// PayBack allow a user to pay back their dues (full or partial)
func (c *Context) PayBack(mr string) {
	m := strings.Split(mr, " ")
	repo := storage.NewRepository(c.DBClient)
	user, err := repo.GetUser(m[1])
	if err != nil {
		fmt.Println("unable to get user from database")
		return
	}

	b, err := strconv.ParseFloat(m[2], 64)
	if err != nil {
		fmt.Println(err)
	}

	if user.Due == 0 || b > user.Due {
		fmt.Println("no dues")
		return
	}

	user.Due = user.Due - b
	user.Balance = user.Balance + b
	err = repo.UpdateUser(*user)
	if err != nil {
		fmt.Println("error updating user", err)
		return

	}

	fmt.Printf("%s(dues: %0.2f)\n", m[1], user.Due)

}
