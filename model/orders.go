package model

type NftOrders struct {
	ID           uint   `json:"id" gorm:"column:id"`
	OrderID      string `json:"order_id" gorm:"column:order_id"`             // 订单号
	GoodsID      uint   `json:"goods_id" gorm:"column:goods_id"`             // 藏品ID
	UserID       uint   `json:"user_id" gorm:"column:user_id"`               // 用户ID
	Num          uint   `json:"num" gorm:"column:num"`                       // 购买数量
	Amount       uint   `json:"amount" gorm:"column:amount"`                 // 订单金额 单位 分
	PayChannelID uint   `json:"pay_channel_id" gorm:"column:pay_channel_id"` // 支付渠道
	PayEndtime   uint   `json:"pay_endtime" gorm:"column:pay_endtime"`       // 支付超时时间
	Status       uint8  `json:"status" gorm:"column:status"`                 // 订单状态 1待支付 2 已支付 3 已取消 4 超时
	DtCreate     uint   `json:"dt_create" gorm:"column:dt_create"`           // 创建时间
	DtUpdate     uint   `json:"dt_update" gorm:"column:dt_update"`           // 更新时间
}

func (m *NftOrders) TableName() string {
	return "nft_orders"
}
