package patternFacade

import (
	"fmt"

	businessLogic "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/facade/business-logic"
)

type WalletFacade struct {
	Account      *businessLogic.Account
	Wallet       *businessLogic.Wallet
	SecurityCode *businessLogic.SecurityCode
	Notification *businessLogic.Notification
	Ledger       *businessLogic.Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	WalletFacacde := &WalletFacade{
		Account:      businessLogic.NewAccount(accountID),
		SecurityCode: businessLogic.NewSecurityCode(code),
		Wallet:       businessLogic.NewWallet(),
		Notification: &businessLogic.Notification{},
		Ledger:       &businessLogic.Ledger{},
	}
	fmt.Println("Account created")
	return WalletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountID, "credit", amount)
	return nil
}
