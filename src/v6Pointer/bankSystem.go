package v6Pointer

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

// 定义string类型打印方式
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance不使用指针无法对数据进行更改

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
