package v6Pointer

import (
	"errors"
	"fmt"
)

// error变量化
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

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

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {

		// errors.New() 创建一个新的error，并带有选择的消息
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

// Balance不使用指针无法对数据进行更改

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/*
总结：
如果写的函数需要改变状态，则必须用指针指向需要修改的值
指针可以是nil，当返回指针时，需要确保其不为nil，避免异常

检查字符串可能会导致测试不稳定，因此可以用一个变量进行代替，更容易测试代码

// errors.New() 创建一个新的error，并带有选择的消息
// t.Fatal(), 如果被调用将停止测试

*/
