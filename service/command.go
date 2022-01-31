package service

import (
	"fmt"
	"os"
	"strings"
)

const (
	user             string = "new user"
	merchant         string = "new merchant"
	transaction      string = "new txn"
	payback          string = "payback"
	update           string = "update merchant"
	discount         string = "report discount"
	dues             string = "report dues"
	usersCreditLimit string = "report users-at-credit-limit"
	usersTotalDues   string = "report total-dues"
	exit             string = "exit"
)

func (c *Context) Command(str string) {
	if strings.HasPrefix(str, user) {
		c.CreateUser(str)
		return
	}
	if strings.HasPrefix(str, merchant) {
		c.CreateMerchant(str)
		return
	}
	if strings.HasPrefix(str, transaction) {
		c.CreateTransaction(str)
		return
	}
	if strings.HasPrefix(str, discount) {
		c.ReportDiscount(str)
		return
	}
	if strings.HasPrefix(str, usersCreditLimit) {
		c.ReportCreditLimit()
		return
	}
	if strings.HasPrefix(str, usersTotalDues) {
		c.ReportTotalDues()
		return
	}
	if strings.HasPrefix(str, dues) {
		c.ReportDuesOfUser(str)
		return
	}
	if strings.HasPrefix(str, update) {
		c.UpdateMerchant(str)
		return
	}
	if strings.HasPrefix(str, payback) {
		c.PayBack(str)
		return
	}
	if strings.HasPrefix(str, exit) {
		os.Stdin.Close()
		return
	}
	fmt.Println("invalid command")
}
