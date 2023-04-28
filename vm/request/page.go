package request

type Page struct {
	PageNo   int `json:"page_no" binding:"required,number"`
	PageSize int `json:"page_size" binding:"required,number"`
}
