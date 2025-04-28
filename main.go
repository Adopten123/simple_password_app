package main

import (
	"fmt"
	"password_app/account"
	"password_app/files"
)

func main() {
	files.WriteFile("Hello", "file.txt")

	var login = promptData("Enter your login: ")
	var password = promptData("Enter your password: ")
	var url = promptData("Enter your url: ")

	var account1, err = account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		panic(err)
	}

	account1.OutputPassword()
	//fmt.Println(account1)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
