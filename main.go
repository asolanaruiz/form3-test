package main

import (
	"context"
	"fmt"
	"github.com/asolanaruiz/form3-accountapi-client"
	"github.com/google/uuid"
)

func main() {
	client := form3.NewClient()
	ctx := context.Background()
	var accountVersion int64 = 1

	account, err := client.Create(createAccount(accountVersion), ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("account from POST: %v\n", *account)
	}

	account, err = client.Fetch(account.Data.ID, ctx)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("account from GET: %v\n", *account)
	}

	ok, err := client.Delete(account.Data.ID, accountVersion, ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("account deleted: %b\n", ok)

	ok, err = client.Delete(account.Data.ID, accountVersion, ctx)
	fmt.Println(err.Error())
}

func createAccount(version int64) form3.Account {
	accountData := form3.AccountData{
		Type: "accounts",
		ID:   uuid.New().String(),
		//ID: "6978c02e-3402-4343-9075-ac00f21c844c",
		//OrganisationID: uuid.New().String(),
		OrganisationID: "0ff772ae-919a-4dc3-9c29-c04f68a803b5",
		Version:        version,
		Attributes: form3.AccountAttributes{
			Country:      "GB",
			BaseCurrency: "GBP",
			BankID:       "11111",
			BankIDCode:   "LLLLL",
			Bic:          "NWBKGB22",
			Name:         []string{"menganito"}}}
	return form3.Account{Data: accountData}
}
