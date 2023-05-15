package request

type Page struct {
	PageNo   int `form:"page_no" binding:"required,number"`
	PageSize int `form:"page_size" binding:"required,number"`
}
