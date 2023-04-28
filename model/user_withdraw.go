package model

type NftUserWithdraw struct {
	ID          uint   `json:"id" gorm:"column:id"`
	UserID      string `json:"user_id" gorm:"column:user_id"`           // 用户编号
	Amount      uint   `json:"amount" gorm:"column:amount"`             // 提现金额
	Bank        string `json:"bank" gorm:"column:bank"`                 // 开户行
	BankName    string `json:"bank_name" gorm:"column:bank_name"`       // 户名
	BankAccount string `json:"bank_account" gorm:"column:bank_account"` // 账号
	BankAddr    string `json:"bank_addr" gorm:"column:bank_addr"`       // 开户地址
	Status      uint8  `json:"status" gorm:"column:status"`             // 状态，1 审核中 2 待打款 3 已打款 4 撤销
	DtCreate    uint   `json:"dt_create" gorm:"column:dt_create"`       // 创建时间
	DtUpdate    uint   `json:"dt_update" gorm:"column:dt_update"`       // 更新时间
}

func (m *NftUserWithdraw) TableName() string {
	return "nft_user_withdraw"
}
