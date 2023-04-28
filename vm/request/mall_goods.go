package request

type MallGoods struct {
	Sort string `json:"sort" binding:"require"`
	Type int    `json:"type" binding:"require,number"`
	Page
}

type MallGoodsList struct {
	Status int `json:"status" binding:"require,number"`
	Page
}

type PlatformGoodsList struct {
	PlatformId int `json:"platform_id" binding:"require,number"`
	Page
}

type MallGoodsAdd struct {
	Name          string `json:"name" binding:"require"`               // 藏品名称
	Image         string `json:"image" binding:"require"`              // 藏品首图
	Detail        string `json:"detail" binding:"require"`             // 藏品介绍
	Price         uint   `json:"price" binding:"require,number"`       // 藏品出售价格
	OriginalPrice uint   `json:"original_price" binding:"-,number"`    // 购买价格
	PlatformID    uint   `json:"platform_id" binding:"require,number"` // 所属平台
	No            string `json:"no" binding:"require"`                 // 编号
	Source        string `json:"source" binding:"-"`                   // 藏品资源
	SourceType    uint8  `json:"source_type" binding:"-"`              // 资源类型 1 图片 2 3D
	Hash          string `json:"hash" binding:"require"`               // HASH地址
}
