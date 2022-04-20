package bankAccount

type BankAccount struct {
	Owner           string
	Agency, Account int
	Balance         float64
}

func (account *BankAccount) Withdraw(amount float64) (string, float64) {
	canWithdraw := amount > 0 && amount <= account.Balance

	if canWithdraw {
		account.Balance -= amount
		return "Withdraw approved", account.Balance
	} else {
		return "Withdraw failed", account.Balance
	}
}

func (account *BankAccount) Deposite(amount float64) (string, float64) {
	canDeposit := amount > 0

	if canDeposit {
		account.Balance += amount
		return "Deposit approved", account.Balance
	} else {
		return "Deposit failed", account.Balance
	}
}

func (account *BankAccount) Transfer(amount float64, target *BankAccount) (string, bool) {
	canTransfer := amount > 0 && account.Balance >= amount

	if canTransfer {
		account.Withdraw(amount)
		target.Deposite(amount)
		return "Transfer approved", true
	} else {
		return "Transfer failed", false
	}
}
