package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string `json:"login"`
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("login required")
	}

	var _, err = url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}

	var tmpAccount = &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		tmpAccount.generatePassword(12)
	}

	return tmpAccount, nil
}

func (account *Account) OutputPassword() {
	color.Cyan(account.login)
	fmt.Println(account.password, account.url)
}

func (account *Account) generatePassword(n int) {
	var password = make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	account.password = string(password)
}
