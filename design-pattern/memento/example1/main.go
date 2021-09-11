package main

import "fmt"

type Memento struct {
	bal int
}

type BankAccount struct {
	bal     int
	changes []*Memento
	current int
}

// func NewBankAccount(amount int) (*BankAccount, *Memento) {
// 	return &BankAccount{bal: amount}, &Memento{amount}
// }

func NewBankAccount(amount int) *BankAccount {
	b := &BankAccount{bal: amount}
	b.changes = append(b.changes, &Memento{bal: amount})
	return b
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.bal += amount
	m := &Memento{b.bal}
	b.changes = append(b.changes, m)
	b.current++
	return m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.bal = m.bal
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.bal = m.bal
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.bal = m.bal
		return m
	}
	return nil
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = %v, current = %v", b.bal, b.current)
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(200)
	ba.Deposit(10000)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println(ba)

	ba.Undo()
	fmt.Println(ba)

	ba.Redo()
	fmt.Println(ba)

	ba.Redo()
	fmt.Println(ba)
}
