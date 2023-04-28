package response

type Login struct {
	UserId   string `json:"user_id"`
	Token    string `json:"token"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
}
