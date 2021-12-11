package account

import (
	"sync"
)

// BankAccount is a representation of an account that can be opened, closed and
// safely accessed concurrently
type Account struct {
	Open bool
	Bal  int64
	m    sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	acc := &Account{
		Open: true,
		Bal:  0,
		m:    sync.Mutex{},
	}
	acc.Lock()
	acc.Bal += amount
	acc.Unlock()
	return acc
}

// Lock locks the internal mutex of the given account so it can be safely accessed
func (a *Account) Lock() {
	a.m.Lock()
}

// Unlock unlocks the internal mutex of the given account so it can be freed up for another caller
func (a *Account) Unlock() {
	a.m.Unlock()
}

// Balance returns the given account's current balance
func (a *Account) Balance() (int64, bool) {
	if !a.Open {
		return 0, false
	}
	return a.Bal, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.Open {
		return amount, false
	}
	a.Lock()
	defer a.Unlock()
	if amount < 0 && (a.Bal+amount) < 0 {
		// If withdrawal would result in negative balance, it must fail
		return 0, false
	}
	a.Bal += amount
	return a.Balance()
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.Open {
		return 0, false
	}
	a.Open = false
	return a.Bal, true
}
