package request

type UserWithdraw struct {
}

type UserWithdrawAdd struct {
	Amount int `json:"amount" binding:"require,number"`
	BankId int `json:"bank_id" binding:"require,number"`
}
