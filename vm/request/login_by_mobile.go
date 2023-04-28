package request

type LoginByMobile struct {
	Mobile string `json:"mobile"  binding:"required,mobile"`
	Code   string `json:"code"  binding:"required,len:6"`
}
