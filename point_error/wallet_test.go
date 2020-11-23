package point

import "testing"

func TestWallet(t *testing.T) {

	testBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		got := w.Balance()
		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		testBalance(t, wallet, want)

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Errorf("no error expect but eccuor")
		}
		testBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw err", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		if err == nil {
			t.Fatal("want err but not")
		}
		if err != MoneyNotEnoughError {
			t.Error("error not expect")
		}
		testBalance(t, wallet, Bitcoin(20))
	})

}
