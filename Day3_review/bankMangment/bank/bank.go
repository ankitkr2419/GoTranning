package bank

import (
	"errors"
	"fmt"

	"github.com/ankitkr1924/bankManagement/account"
)

var accounts map[string]account.Account

func init() {
	accounts = make(map[string]account.Account)
}

type Bank interface {
	CreateAccount(name string) (id string)
	DeleteAccount(id string)
	DepositeMoney(id string, amount float64)
	WithDrawMoney(id string, amount float64)
	GetAccountHistory(id string)
}

func NewBank() *bankInfo {
	return &bankInfo{}
}

type bankInfo struct {
}

func (b *bankInfo) CreateAccount(name string) (id string) {
	acc, id := account.NewAccount(name)
	accounts[id] = acc
	return id
}

func (b *bankInfo) DeleteAccount(id string) {
	_, ok := accounts[id]
	if !ok {
		fmt.Println("Please Enter Valid Id")
		return
	}
	delete(accounts, id)
}

func (b *bankInfo) DepositeMoney(id string, amount float64) {
	acc, err := getAccount(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = acc.Deposite(amount)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (b *bankInfo) WithDrawMoney(id string, amount float64) {
	acc, err := getAccount(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = acc.Withdraw(amount)
	if err != nil {
		fmt.Println(err)
	}
}

func (b *bankInfo) GetAccountHistory(id string) {
	acc, err := getAccount(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	acc.History()
}

func getAccount(id string) (acc account.Account, err error) {
	acc, ok := accounts[id]
	if !ok {
		err := errors.New("no account found")
		return nil, err
	}
	return acc, nil
}
