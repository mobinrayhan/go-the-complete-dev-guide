package bankaccount

type BankAccount struct {
	Owner   string
	Balance *float64
}

func (ba *BankAccount) Deposit(amount float64) {
	*ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) {
	if *ba.Balance < amount {
		panic("Your Withdraw Money is lower than you bank account!")
	}
	*ba.Balance -= amount
}

func (ba *BankAccount) MainBalance() float64 {
	return *ba.Balance
}
