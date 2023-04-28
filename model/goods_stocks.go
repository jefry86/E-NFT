package model

type NftGoodsStocks struct {
	ID        uint   `json:"id" gorm:"column:id"`
	GoodsID   uint   `json:"goods_id" gorm:"column:goods_id"`     // 藏品ID
	GoodsNo   string `json:"goods_no" gorm:"column:goods_no"`     // 藏品编号
	GoodsHash string `json:"goods_hash" gorm:"column:goods_hash"` // 藏品HASH地址
	Status    uint8  `json:"status" gorm:"column:status"`         // 是否可用 0不可用 1 可用
	UserID    string `json:"user_id" gorm:"column:user_id"`       // 用户编号
	DtCreate  uint   `json:"dt_create" gorm:"column:dt_create"`   // 创建时间
	DtUpdate  uint   `json:"dt_update" gorm:"column:dt_update"`   // 更新时间
}

func (m *NftGoodsStocks) TableName() string {
	return "nft_goods_stocks"
}
