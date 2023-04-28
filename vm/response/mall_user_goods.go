package response

type MallUserGoodsList struct {
	Id           uint   `json:"id"`
	GoodsId      uint   `json:"goods_id"`
	GoodsName    string `json:"goods_name"`
	GoodsImage   string `json:"goods_image"`
	GoodsType    uint8  `json:"goods_type"`
	GoodsLabel   string `json:"goods_label"`
	GoodsNo      string `json:"goods_no"`
	PlatformId   uint   `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	PlatformLogo string `json:"platform_logo"`
	DateTime     string `json:"date_time"`
}
type MallUserGoodsRes struct {
	List []MallUserGoodsList `json:"list"`
	Page
}

type MallUserGoodsInfo struct {
	Id           uint   `json:"id"`
	GoodsId      uint   `json:"goods_id"`
	GoodsName    string `json:"goods_name"`
	GoodsImage   string `json:"goods_image"`
	GoodsType    uint8  `json:"goods_type"`
	GoodsLabel   string `json:"goods_label"`
	GoodsNo      string `json:"goods_no"`
	GoodsDetail  string `json:"goods_detail"`
	GoodsHash    string `json:"goods_hash"`
	PlatformId   uint   `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	PlatformLogo string `json:"platform_logo"`
	DateTime     string `json:"date_time"`
}
