package response

import "time"

type UserWithdrawInfo struct {
	Amount   uint      `json:"amount"` // 提现金额
	Bank     string    `json:"bank"`   // 开户行
	Status   uint8     `json:"status"` // 状态，1 审核中 2 待打款 3 已打款 4 撤销
	DateTime time.Time `json:"date_time"`
}

type UserWithdrawListRes struct {
	List []UserWithdrawInfo
	Page
}
