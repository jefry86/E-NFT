package request

type SendSMS struct {
	Mobile string `form:"mobile" binding:"required,mobile"`
	T      string `form:"type" binding:"required,number"`
}
