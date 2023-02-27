package wallet

import "fmt"
import "errors"

type Bitcoin float64
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	return w.balance
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

func (w *Wallet) Withdraw(b Bitcoin) (e error) {
	if b > w.balance {
		return errors.New("oh no")
	}
	w.balance -= b
	return nil
}
