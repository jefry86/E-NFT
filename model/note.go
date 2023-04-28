package model

type NftNote struct {
	ID       uint   `json:"id" gorm:"column:id"`
	Title    string `json:"title" gorm:"column:title"`         // 标题
	Intro    string `json:"intro" gorm:"column:intro"`         // 简介
	Content  string `json:"content" gorm:"column:content"`     // 内容
	Status   uint8  `json:"status" gorm:"column:status"`       // 是否展示 0 不展示 1 展示
	Type     uint8  `json:"type" gorm:"column:type"`           // 类型 1 自营平台2 寄售平台
	DtCreate uint   `json:"dt_create" gorm:"column:dt_create"` // 创建时间
	DtUpdate uint   `json:"dt_update" gorm:"column:dt_update"` // 更新时间
}

func (m *NftNote) TableName() string {
	return "nft_note"
}
