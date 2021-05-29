package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))

		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(20))
		_, err := w.Withdraw(Bitcoin(10))

		assertBalance(t, w, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		w := Wallet{}
		w.Deposit(startingBalance)

		_, err := w.Withdraw(Bitcoin(100))
		assertBalance(t, w, startingBalance)
		assertError(t, ErrInsufficientFunds, err)
	})
}

func assertBalance(t testing.TB, w Wallet, expected Bitcoin) {
	t.Helper()
	actual := w.Balance()

	if actual != expected {
		t.Errorf("expected %s; actual %s", expected, actual)
	}
}

func assertNoError(t testing.TB, actual error) {
	t.Helper()

	if actual != nil {
		t.Fatal("got an error but didn’t want one")
	}
}

func assertError(t testing.TB, expected error, actual error) {
	t.Helper()

	if actual == nil {
		t.Fatal("expected an error but got nil")
	}

	if actual != expected {
		t.Errorf("expected %q; actual %q", expected, actual)
	}
}
