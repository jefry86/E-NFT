package response

type UserBankInfo struct {
	Bank        string `json:"bank"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	BankAddr    string `json:"bank_addr"`
}
