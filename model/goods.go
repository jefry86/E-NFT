package model

type NftGoods struct {
	ID                   uint   `json:"id" gorm:"column:id"`
	No                   string `json:"no" gorm:"column:no"`                                       // 批次
	Name                 string `json:"name" gorm:"column:name"`                                   // 藏品名称
	Image                string `json:"image" gorm:"column:image"`                                 // 藏品首图 400*400
	Type                 int8   `json:"type" gorm:"column:type"`                                   // 藏品类型 1普通藏品 2盲盒
	Detail               string `json:"detail" gorm:"column:detail"`                               // 藏品详情图
	Label                string `json:"label" gorm:"column:label"`                                 // 标签 ， |分割
	Price                uint   `json:"price" gorm:"column:price"`                                 // 价格，单位 分
	AppleProductID       string `json:"apple_product_id" gorm:"column:apple_product_id"`           // 苹果IAP 商品ID
	IsShow               uint8  `json:"is_show" gorm:"column:is_show"`                             // 是否显示
	HasSaled             uint8  `json:"has_saled" gorm:"column:has_saled"`                         // 是否售罄
	SaleStarttime        uint   `json:"sale_starttime" gorm:"column:sale_starttime"`               // 销售时间
	SaleEndtime          uint   `json:"sale_endtime" gorm:"column:sale_endtime"`                   // 销售结束时间
	AppointmentStarttime uint   `json:"appointment_starttime" gorm:"column:appointment_starttime"` // 预约时间
	AppointmentEndtime   uint   `json:"appointment_endtime" gorm:"column:appointment_endtime"`     // 预约结束时间
	Stock                uint   `json:"stock" gorm:"column:stock"`                                 // 库存
	SalesVolume          uint   `json:"sales_volume" gorm:"column:sales_volume"`                   // 销量
	PublisherID          uint   `json:"publisher_id" gorm:"column:publisher_id"`                   // 发行商
	Source               string `json:"source" gorm:"column:source"`                               // 资源地址
	SourceType           uint8  `json:"source_type" gorm:"column:source_type"`                     // 资源类型 1 图片 2 3D
	HasCast              uint8  `json:"has_cast" gorm:"column:has_cast"`                           // 是否已铸造
	DtCreate             uint   `json:"dt_create" gorm:"column:dt_create"`                         // 创建时间
	DtUpdate             uint   `json:"dt_update" gorm:"column:dt_update"`                         // 更新时间
}

func (m *NftGoods) TableName() string {
	return "nft_goods"
}
