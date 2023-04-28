package response

import "time"

type MallOrderApply struct {
	OrderId string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

type MallOrderInfo struct {
	OrderId    string    `json:"order_id"`
	Amount     uint      `json:"amount"`
	Num        uint      `json:"num"`
	GoodsId    uint      `json:"goods_id"`
	GoodsName  string    `json:"goods_name"`
	GoodsImage string    `json:"goods_image"`
	GoodsType  uint8     `json:"goods_type"`
	GoodsLabel string    `json:"goods_label"`
	GoodsNo    string    `json:"goods_no"`
	DateTime   time.Time `json:"date_time"`
	PayEndTime uint      `json:"pay_end_time"`
	Status     uint8     `json:"status"`
}

type MallOrderList struct {
	OrderId    string    `json:"order_id"`
	Amount     uint      `json:"amount"`
	Num        uint      `json:"num"`
	GoodsId    uint      `json:"goods_id"`
	GoodsName  string    `json:"goods_name"`
	GoodsImage string    `json:"goods_image"`
	GoodsType  uint8     `json:"goods_type"`
	GoodsLabel string    `json:"goods_label"`
	GoodsNo    string    `json:"goods_no"`
	DateTime   time.Time `json:"date_time"`
	Status     uint8     `json:"status"`
}

type MallOrderListRes struct {
	List []MallOrderList `json:"list"`
	Page
}
