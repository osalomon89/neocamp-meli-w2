package wallet

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Parallel()
	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Deposit(15)

		result := wallet.Balance()
		expected := 25

		if expected != result {
			t.Errorf("result %d and expected %d", result, expected)
		}
	})

	t.Run(("Withdraw"), func(t *testing.T) {
		wallet := Wallet{
			balance: 100,
		}

		err := wallet.Withdraw(30)
		if err != nil {
			t.Errorf("error not expected %s", err.Error())
		}

		result := wallet.Balance()
		expected := 70

		if expected != result {
			t.Errorf("result %d and expected %d", result, expected)
		}
	})

	t.Run(("Withdraw Error"), func(t *testing.T) {
		wallet := Wallet{
			balance: 20,
		}

		err := wallet.Withdraw(30)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		expected := "not enough money"

		if err.Error() != expected {
			t.Errorf("result '%s' and expected '%s'", err, expected)
		}
	})
}

func TestWalletTD(t *testing.T) {
	tests := []struct {
		name        string
		wallet      Wallet
		isDeposit   bool
		amount      int
		want        int
		wantedError error
	}{
		{
			name: "Balance OK: Deposit",
			wallet: Wallet{
				balance: 20,
			},
			isDeposit:   true,
			amount:      25,
			want:        45,
			wantedError: nil,
		},
		{
			name: "Balance OK: Withdraw",
			wallet: Wallet{
				balance: 100,
			},
			isDeposit:   false,
			amount:      25,
			want:        75,
			wantedError: nil,
		},
		{
			name: "Balance Error: Withdraw",
			wallet: Wallet{
				balance: 20,
			},
			isDeposit:   false,
			amount:      25,
			wantedError: errors.New("not enough money"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isDeposit {
				tt.wallet.Deposit(tt.amount)
			} else {
				err := tt.wallet.Withdraw(tt.amount)
				if tt.wantedError != nil {
					if err.Error() != tt.wantedError.Error() {
						t.Errorf("result '%s' and expected '%s'", err, tt.wantedError)
					}
					return
				}

				if err != nil {
					t.Fatal("unexpected error")
				}
			}

			result := tt.wallet.Balance()

			if tt.want != result {
				t.Errorf("result %d and expected %d", result, tt.want)
			}
		})
	}
}

/*func TestWalletWithdraw(t *testing.T) {
	wallet := Wallet{
		balance: 100,
	}

	err := wallet.Withdraw(30)
	if err != nil {
		t.Errorf("error not expected %s", err.Error())
	}

	result := wallet.Balance()
	expected := 70

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}

func TestWalletWithdrawError(t *testing.T) {
	wallet := Wallet{
		balance: 20,
	}

	err := wallet.Withdraw(30)
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	expected := "not enough money"

	if err.Error() != expected {
		t.Errorf("result '%s' and expected '%s'", err, expected)
	}
}*/
