package main

import (
	"fmt"
	"strings"
	"test/test_app_4/account"
	"test/test_app_4/encrypter"
	// "test/test_app_4/cloud"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"test/test_app_4/files"
	"test/test_app_4/output"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}
var menuVariants = []string{
	"Choose an action",
	"1- Create;",
	"2- Find by URL;",
	"3- Find by login;",
	"4- Delete;",
	"5- Exit.",
	"Your choise",
}

func counterFunc() func() {
	i := 0
	return func() {
		i++
		fmt.Printf("The menu was opened %v times\n", i)
	}
}

func main() {
	fmt.Println("__ Password Manager __")
	err := godotenv.Load()
	if err != nil {
		output.PrintError(err)
	}
	counter := counterFunc()
	vault, _ := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
	// vault, _ := account.NewVault(cloud.NewCloudDB("https://google.com"))
	for {
		counter()
		choise := promptData(menuVariants...)
		if choise == "5" {
			return
		}
		menuFunc := menu[choise]
		if menuFunc == nil {
			output.PrintError("Empty or wrong input")
		} else {
			menuFunc(vault)
		}
	}
}

func promptData(prompt ...string) string {
	var userInput string
	for pos, elem := range prompt {
		if pos == len(prompt)-1 {
			fmt.Printf("%v: ", elem)
		} else {
			fmt.Println(elem)
		}
	}
	fmt.Scanln(&userInput)
	return userInput
}

func createAccount(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	vault.AddAccount(myAccount)
}

func findAccountByURL(vault *account.VaultWithDB) {
	loginURL := promptData("Enter URL to search")
	accounts, _ := vault.FindAccount(loginURL, CheckUrl)
	outputResultsOfSearch(accounts)

}
func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData("Enter login to search")
	accounts, _ := vault.FindAccount(login, CheckLogin)
	outputResultsOfSearch(accounts)
}
func CheckUrl(acc account.Account, urlInput string) bool {
	return strings.Contains(acc.Url, urlInput)
}
func CheckLogin(acc account.Account, loginInput string) bool {
	return strings.Contains(acc.Login, loginInput)
}
func outputResultsOfSearch(accounts []account.Account) {
	if len(accounts) == 0 {
		output.PrintError("Account not found")
	}
	for _, acc := range accounts {
		color.HiGreen("Succsess search: ")
		acc.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	loginURL := promptData("Enter URL to delete")
	err := vault.DeleteAccount(loginURL)
	if err != nil {
		output.PrintError(err)
	} else {
		color.Green("Succsess deleted")
	}
}
