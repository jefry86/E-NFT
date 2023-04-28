package request

type NotifyPay struct {
	OrderId    string `json:"order_id"`
	Amount     uint   `json:"amount"`
	PayTime    string `json:"pay_time"`
	PayChannel string `json:"pay_channel"`
}
