package wallet

import "errors"

type Wallet struct {
	balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() int {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount int) error {
	if amount > wallet.balance {
		return errors.New("not enough money")
	}

	wallet.balance -= amount
	return nil
}
