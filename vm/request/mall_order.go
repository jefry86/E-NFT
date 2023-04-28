package request

type MallOrderApply struct {
	GoodsId uint `json:"goods_id" binding:"require,number"`
	Num     uint `json:"num" binding:"require,number"`
}

type MallOrderList struct {
	Status int `json:"status" binding:"require,number"`
	Page
}
