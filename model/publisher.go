package model

type NftPublisher struct {
	ID       uint   `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`           // 名称
	Image    string `json:"image" gorm:"column:image"`         // 头像
	Status   uint8  `json:"status" gorm:"column:status"`       // 状态 0 禁用 1 启用
	DtCreate uint   `json:"dt_create" gorm:"column:dt_create"` // 创建时间
	DtUpdate uint   `json:"dt_update" gorm:"column:dt_update"` // 更新时间
}

func (m *NftPublisher) TableName() string {
	return "nft_publisher"
}
