package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, expected Bitcoin) {
		t.Helper()
		actual := w.Balance()

		if actual != expected {
			t.Errorf("expected %s; actual %s", expected, actual)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(20))

		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{}
		w.Deposit(startingBalance)

		_, err := w.Withdraw(Bitcoin(100))
		assertBalance(t, w, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn’t get one")
		}
	})
}
