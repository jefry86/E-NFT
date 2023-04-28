package response

type MallGoodsList struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Type         uint8  `json:"type"`
	Label        string `json:"label"`
	No           string `json:"no"`
	PlatformId   uint   `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	PlatformLogo string `json:"platform_logo"`
	DateTime     string `json:"date_time"`
}

type MallGoodsRes struct {
	List []MallGoodsList `json:"list"`
	Page
}

type MallGoodsInfo struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Type         uint8  `json:"type"`
	Label        string `json:"label"`
	No           string `json:"no"`
	Detail       string `json:"detail"`
	Hash         string `json:"hash"`
	PlatformId   uint   `json:"platform_id"`
	PlatformName string `json:"platform_name"`
	PlatformLogo string `json:"platform_logo"`
	DateTime     string `json:"date_time"`
}
