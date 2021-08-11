package businessLogic

import "fmt"

type Account struct {
	Name string
}

func newAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
