package response

type MallPlatformList struct {
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Site      string `json:"site"`
	HasOnline uint8  `json:"has_online"`
}

type MallPlatformListRes struct {
	List []MallPlatformList `json:"list"`
	Page
}
