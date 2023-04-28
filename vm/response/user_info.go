package response

type UserInfo struct {
	UserId     string `json:"user_id"`
	Nickname   string `json:"nickname"`
	Mobile     string `json:"mobile"`
	Avatar     string `json:"avatar"`
	WalletAddr string `json:"wallet_addr"`
}
