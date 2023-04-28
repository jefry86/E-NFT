package response

type PlatformUser struct {
	UserId     string `json:"user_id"`
	WalletHash string `json:"wallet_hash"`
}

type PlatformTransfer struct {
	State  int `json:"state"`
	TaskId int `json:"task_id"`
}

type PlatformSales struct {
	CanSale int `json:"can_sale"`
}

type PlatformGoodsListRes struct {
	Total    int                 `json:"total"`
	PageSize int                 `json:"page_size"`
	List     []PlatformGoodsList `json:"list"`
}

type PlatformGoodsList struct {
	GoodsId      int    `json:"goods_id"`
	GoodsHash    string `json:"goods_hash"`
	GoodsName    string `json:"goods_name"`
	GoodsNo      string `json:"goods_no"`
	Image        string `json:"image"`
	DetailImage  string `json:"detail_image"`
	Creator      string `json:"creator"`
	Platform     string `json:"platform"`
	DetailText   string `json:"detail_text"`
	Source       string `json:"source"`
	GoodsType    int    `json:"goods_type"`
	PurchaseTime string `json:"purchase_time"`
	SaleTime     string `json:"sale_time"`
	Price        string `json:"price"`
	Extend       string `json:"extend"`
}
