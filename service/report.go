package service

import (
	"fmt"
	"pay-later-service/storage"
	"strings"
)

// ReportCreditLimit reports which users have reached their credit limit
func (c *Context) ReportCreditLimit() {
	repo := storage.NewRepository(c.DBClient)

	user, err := repo.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range user {
		fmt.Println(u.Name)
	}
}

// ReportDiscount reports how much discount we received from a merchant till date
func (c *Context) ReportDiscount(mr string) {
	m := strings.Split(mr, " ")
	repo := storage.NewRepository(c.DBClient)
	user, err := repo.GetMerchant(m[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.TotalDiscount)
}

// ReportTotalDues reports dues for all users
func (c *Context) ReportTotalDues() {
	repo := storage.NewRepository(c.DBClient)

	user, err := repo.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range user {
		if v.Due != 0 {
			fmt.Println(v.Name, ":", v.Due)
		}
	}
}

// ReportDueOfUser reports dues for a user so far
func (c *Context) ReportDuesOfUser(mr string) {
	m := strings.Split(mr, " ")
	repo := storage.NewRepository(c.DBClient)
	user, err := repo.GetUser(m[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Name, ":", user.Due)
}
