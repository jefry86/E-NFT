package model

type NftMangheConf struct {
	ID       uint   `json:"id" gorm:"column:id"`
	MangheID uint   `json:"manghe_id" gorm:"column:manghe_id"` // 盲盒ID
	GoodsID  uint   `json:"goods_id" gorm:"column:goods_id"`   // 藏品ID
	Rate     uint   `json:"rate" gorm:"column:rate"`           // 概率
	Level    string `json:"level" gorm:"column:level"`         // 等级 SSSR SSR SR R
	Status   uint8  `json:"status" gorm:"column:status"`       // 是否可用
	DtCreate uint   `json:"dt_create" gorm:"column:dt_create"` // 创建时间
	DtUpdate uint   `json:"dt_update" gorm:"column:dt_update"` // 更新时间
}

func (m *NftMangheConf) TableName() string {
	return "nft_manghe_conf"
}
