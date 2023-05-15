package request

type MallGoods struct {
	Sort       string `form:"sort"`
	PlatformId string `form:"platform_id"`
	Type       int    `form:"type" binding:"required,number"`
	Page
}

type MallGoodsList struct {
	Status int `form:"status" binding:"number"`
	Page
}

type MallGoodsSearch struct {
	Sort       string `form:"sort"`
	PlatformId string `form:"platform_id"`
	Type       int    `form:"type" binding:"number"`
	Keyword    string `form:"keyword" binding:"required"`
	Page
}

type PlatformGoodsList struct {
	PlatformId int `json:"platform_id" binding:"require,number"`
	Page
}

type MallGoodsAdd struct {
	Name          string `form:"name" json:"name" binding:"required"`                      // 藏品名称
	Image         string `form:"image" json:"image" binding:"required"`                    // 藏品首图
	Detail        string `form:"detail" json:"detail"`                                     // 藏品介绍
	Price         uint   `form:"price" json:"price" binding:"required,number"`             // 藏品出售价格
	OriginalPrice uint   `form:"original_price" json:"original_price" binding:"number"`    // 购买价格
	PlatformID    uint   `form:"platform_id" json:"platform_id" binding:"required,number"` // 所属平台
	No            string `form:"no" json:"no"`                                             // 编号
	Source        string `form:"source" json:"source"`                                     // 藏品资源
	SourceType    uint8  `form:"source_type" json:"source_type"`                           // 资源类型 1 图片 2 3D
	Hash          string `form:"hash" json:"hash"`                                         // HASH地址
}
