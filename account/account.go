package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-")
type Account  struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Url       string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		// return  nil, err
	}

	if len(login) == 0 {
		return nil, errors.New("invalid login")
	}

	acc := &Account{
		Login: login,
		Password: password,
		Url: urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),	
	}

	if acc.Password == ""{
		acc.GeneratePassword(12)
	}

	return acc, nil
}

func (acc *Account) OutputData (){
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) GeneratePassword ( lenght int)  {
	password := make([]rune, lenght)
	for pos := range password {
		password[pos] = letters[rand.IntN(len(letters))]
	}
	acc.Password = string(password)
}
