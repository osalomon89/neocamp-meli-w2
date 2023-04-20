package wallet_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamcubation/neocamp-meli/testing/wallet"
)

func TestWalletDeposit(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 10,
	}

	wallet.Deposit(15)

	result := wallet.Balance
	expected := 25

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}

func TestWalletWithdraw(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 100,
	}

	err := wallet.Withdraw(30)
	if err != nil {
		t.Errorf("error not expected %s", err.Error())
	}

	result := wallet.Balance
	expected := 70

	if expected != result {
		t.Errorf("result %d and expected %d", result, expected)
	}
}

func TestWalletWithdrawError(t *testing.T) {
	wallet := wallet.Wallet{
		Balance: 20,
	}

	err := wallet.Withdraw(30)
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}

	expected := "not enough money"

	if err.Error() != expected {
		t.Errorf("result '%s' and expected '%s'", err, expected)
	}
}

func TestWallet(t *testing.T) {
	t.Parallel()
	t.Run("Balance", func(t *testing.T) {
		wallet := wallet.Wallet{
			Balance: 10,
		}

		wallet.Deposit(15)

		result := wallet.Balance
		expected := 25

		if expected != result {
			t.Errorf("result %d and expected %d", result, expected)
		}
	})

	t.Run(("Withdraw"), func(t *testing.T) {
		wallet := wallet.Wallet{
			Balance: 100,
		}

		err := wallet.Withdraw(30)
		if err != nil {
			t.Errorf("error not expected %s", err.Error())
		}

		result := wallet.Balance
		expected := 70

		if expected != result {
			t.Errorf("result %d and expected %d", result, expected)
		}
	})

	t.Run(("Withdraw Error"), func(t *testing.T) {
		wallet := wallet.Wallet{
			Balance: 20,
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

func TestWalletWithdrawTDT(t *testing.T) {
	tests := []struct {
		name        string
		wallet      wallet.Wallet
		amount      int
		want        int
		wantedError error
	}{
		{
			name: "Balance OK: Withdraw",
			wallet: wallet.Wallet{
				Balance: 100,
			},
			amount:      25,
			want:        75,
			wantedError: nil,
		},
		{
			name: "Balance Error: Withdraw",
			wallet: wallet.Wallet{
				Balance: 20,
			},
			amount:      45,
			wantedError: errors.New("not enough money"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wallet.Withdraw(tt.amount)
			if tt.wantedError != nil {
				assert.NotNil(t, err, "wanted an error but didn't get one")
				//assert.NotEqual(t, err, nil, "wanted an error but didn't get one")
				/*if err == nil {
					t.Error("wanted an error but didn't get one")
				}*/

				assert.Equal(t, err, tt.wantedError, "they should be equal")
				/*if err.Error() != tt.wantedError.Error() {
					t.Errorf("result '%s' and expected '%s'", err, tt.wantedError)
				}*/
				return
			}

			assert.Nil(t, err, "unexpected error")
			assert.Equal(t, tt.wallet.Balance, tt.want, "they should be equal")
			/*if err != nil {
				t.Fatal("unexpected error")
			}

			result := tt.wallet.Balance

			if tt.want != result {
				t.Errorf("result %d and expected %d", result, tt.want)
			}*/
		})
	}
}
