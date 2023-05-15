package request

type LoginByMobile struct {
	Mobile string `form:"mobile" json:"mobile"  binding:"required,mobile"`
	Code   string `form:"code" json:"code"  binding:"required,number,len=6"`
}
