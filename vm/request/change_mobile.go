package request

type ChangeMobile struct {
	Mobile string `json:"mobile" binding:"-"`
	Code   string `json:"code" binding:"require,len:6"`
	Step   int    `json:"step" binding:"-"`
}
