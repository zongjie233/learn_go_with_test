package v6Pointer

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance不使用指针无法对数据进行更改

func (w *Wallet) Balance() int {
	return w.balance
}
