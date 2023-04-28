package model

type NftUserMsg struct {
	ID         uint   `json:"id" gorm:"column:id"`
	Title      string `json:"title" gorm:"column:title"`               // 标题
	Content    string `json:"content" gorm:"column:content"`           // 消息内容
	UserID     string `json:"user_id" gorm:"column:user_id"`           // 用户编号
	FromUserID string `json:"from_user_id" gorm:"column:from_user_id"` // 来自
	HasRead    int8   `json:"has_read" gorm:"column:has_read"`         // 是否一度
	DtCreate   uint   `json:"dt_create" gorm:"column:dt_create"`       // 创建时间
	DtUpdate   uint   `json:"dt_update" gorm:"column:dt_update"`       // 更新时间
}

func (m *NftUserMsg) TableName() string {
	return "nft_user_msg"
}
