package main

import (
	"context"
	"fmt"
	"github.com/asolanaruiz/form3-accountapi-client"
  "github.com/google/uuid"
)

func main() {
  client := form3.NewClient().WithUserAgent("asolanaruiz")
	ctx := context.Background()

	account, err := client.Create(createAccount(), ctx)

	if err != nil {
		panic(err)
	}

	fmt.Printf("account: %v", *account)
}


func createAccount() form3.Account {
  accountData := form3.AccountData{
    Type: "accounts",
    ID: uuid.New().String(),
    OrganisationID: uuid.New().String(),
    Attributes: form3.AccountAttributes{
      Country: "GB",
      BaseCurrency: "GBP",
      BankID: "400300",
	    BankIDCode: "ABDSC",
      Bic: "NWBKGB22",
      Name: []string{"fulanito"}}}
  return form3.Account{Data: accountData}
}
