package model

type NftUserGoods struct {
	ID             uint   `json:"id" gorm:"column:id"`
	GoodsID        uint   `json:"goods_id" gorm:"column:goods_id"`                 // 藏品ID
	GoodsName      string `json:"goods_name" gorm:"column:goods_name"`             // 藏品名称
	GoodsNo        string `json:"goods_no" gorm:"column:goods_no"`                 // 藏品编号
	UserID         string `json:"user_id" gorm:"column:user_id"`                   // 用户编号
	GoodsType      int8   `json:"goods_type" gorm:"column:goods_type"`             // 藏品类型 1 普通藏品 2 盲盒
	GoodsHash      string `json:"goods_hash" gorm:"column:goods_hash"`             // 藏品HASh
	FromWalletAddr string `json:"from_wallet_addr" gorm:"column:from_wallet_addr"` // 来源用户钱包地址，空来源系统发行
	Status         uint8  `json:"status" gorm:"column:status"`                     // 是否可用
	DtCreate       uint   `json:"dt_create" gorm:"column:dt_create"`               // 创建时间
	DtUpdate       uint   `json:"dt_update" gorm:"column:dt_update"`               // 更新时间
}

func (m *NftUserGoods) TableName() string {
	return "nft_user_goods"
}
