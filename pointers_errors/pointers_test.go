package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit to Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)

	})
	t.Run("Withdraw from the Wallet", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(10),
		}

		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})

}

func assertBalance(t testing.TB, wallet Wallet, balance Bitcoin) {
	t.Helper()
	if wallet.Balance() != balance {
		t.Errorf("got %s and wanted %s", wallet.Balance(), balance)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("Expected nil error but got an Error")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but no error")
	}

	if got != want {
		t.Errorf("Got %q, Expected %q", got, want)
	}
}
