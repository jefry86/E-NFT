package response

import "time"

type UserBalanceList struct {
	Amount   int       `json:"amount"`
	Balance  uint      `json:"balance"`
	Info     string    `json:"info"`
	Type     uint8     `json:"type"`
	DateTime time.Time `json:"date_time"`
}

type UserBalanceListRes struct {
	List []UserBalanceList `json:"list"`
	Page
}
