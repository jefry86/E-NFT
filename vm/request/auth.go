package request

type Auth struct {
	RealName string `json:"real_name" binding:"required,chinese"`
	CardNo   string `json:"card_no" binding:"required,cardno"`
}
