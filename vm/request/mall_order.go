package request

type MallOrderApply struct {
	GoodsId uint `json:"goods_id" binding:"require,number"`
	Num     uint `json:"num" binding:"require,number"`
}

type MallOrderList struct {
	T int `json:"type" binding:"require,number"`
	Page
}
