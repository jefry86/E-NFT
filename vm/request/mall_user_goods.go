package request

type MallUserGoods struct {
	Type int `json:"type" binding:"require,number"`
	Page
}
