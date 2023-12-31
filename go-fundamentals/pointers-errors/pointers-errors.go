package pointerserrors

import "fmt"

// int is fine in terms of the way it works but it's not descriptive.
// we're making a new type and we can declare methods on them
type Bitcoin int

// Stringer -> Stringer is implemented by any value that has a String method
// lets you define how your type is printed when used with the %s format string in prints.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// At some point you may wish to use structs to manage state, exposing methods to let users change the state in a way that you can control.
type Wallet struct {
	// a lowercase symbol then it is private outside the package it's defined in.
	balance Bitcoin
}

// func (receiverName ReceiverType) MethodName(args)
// we instead take a pointer to that wallet so that we can change the original values within it.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// In Go, when you call a function or a method the arguments are copied without pointer.
func (w *Wallet) Balance() Bitcoin {
	// "automatically dereferenced" the same of "return (*w).balance"
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
