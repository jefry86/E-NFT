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

type AddByApi struct {
	PlatformId int `json:"platform_id" binding:"require,number"`
	Page
}
