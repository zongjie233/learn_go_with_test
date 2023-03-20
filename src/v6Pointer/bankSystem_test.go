package v6Pointer

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)

		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			// Fatal(), 如果被调用将停止测试
			t.Fatal("didnt get an error but wanted one")
		}

		if got != want {
			t.Errorf("got '%s', want'%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})

}
