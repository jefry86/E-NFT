package model

type NftMallPlatform struct {
	ID           uint   `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`                     // 平台名称
	Logo         string `json:"logo" gorm:"column:logo"`                     // 平台logo
	Site         string `json:"site" gorm:"column:site"`                     // 平台web地址
	ApiUser      string `json:"api_user" gorm:"column:api_user"`             // 登记api
	ApiGoodsList string `json:"api_goods_list" gorm:"column:api_goods_list"` // 藏品列表API
	ApiTransfer  string `json:"api_transfer" gorm:"column:api_transfer"`     // 藏品转移API
	ApiSales     string `json:"api_sales" gorm:"column:api_sales"`           // 藏品上架销售查询API
	Status       uint8  `json:"status" gorm:"column:status"`                 // 是否可用 1可用 0 不可用
	DtCreate     uint   `json:"dt_create" gorm:"column:dt_create"`           // 创建时间
	DtUpdate     uint   `json:"dt_update" gorm:"column:dt_update"`           // 跟新时间
}

func (m *NftMallPlatform) TableName() string {
	return "nft_mall_platform"
}
