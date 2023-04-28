package response

type Page struct {
	Total int64 `json:"total"`
	Size  int   `json:"size"`
}
