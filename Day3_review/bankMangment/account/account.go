package account

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var accountHistory map[string][]transation

func init() {
	accountHistory = make(map[string][]transation)
}

type TransationType string

const (
	TransationTypeCredit TransationType = "credit"
	TransationTypeDebit  TransationType = "debit"
)

type transation struct {
	account        AccountInfo
	time           time.Time
	amount         float64
	transationType TransationType
}

type Account interface {
	Deposite(amount float64) (err error)
	Withdraw(amount float64) (err error)
	History()
}

type AccountInfo struct {
	id      string
	name    string
	balance float64
}

func (a *AccountInfo) modify(tx TransationType, amount float64) {
	switch tx {
	case TransationTypeCredit:
		a.balance = a.balance - amount
	case TransationTypeDebit:
		a.balance = a.balance + amount
	}
}

func (a *AccountInfo) Deposite(amount float64) (err error) {
	tx, err := newTransation(TransationTypeDebit, amount, a)
	if err != nil {
		return err
	}

	accountHistory[a.id] = append(accountHistory[a.id], tx)
	return
}

func (a *AccountInfo) Withdraw(amount float64) (err error) {
	tx, err := newTransation(TransationTypeCredit, amount, a)
	if err != nil {
		return err
	}

	accountHistory[a.id] = append(accountHistory[a.id], tx)
	return
}

func (a *AccountInfo) History() {
	fmt.Printf("|%-20s|%20s|%-20s|%-20s\n", "Time", "TransactionType", "Amount", "Account Balance")
	for _, tran := range accountHistory[a.id] {
		fmt.Printf("|%-20v|%20v|%-20v|%-20v\n", tran.time.Local().Format("Sun Jan 16 16:53:05"), tran.transationType, tran.amount, tran.account.balance)
	}
}

func newTransation(txType TransationType, amount float64, account *AccountInfo) (tx transation, err error) {
	fmt.Println("Strating New Transation")
	if account.balance-amount < 0 && txType == TransationTypeCredit {
		err = errors.New("dont have enough account balance")
		return transation{}, err
	}
	account.modify(txType, amount)

	return transation{*account, time.Now(), amount, txType}, err
}

func NewAccount(name string) (acc Account, id string) {
	id_gen := uuid.NewString()
	acc = &AccountInfo{
		id:      id_gen,
		name:    name,
		balance: 0,
	}
	return acc, id_gen
}
