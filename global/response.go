package global

type Code int

const (
	OK               Code = 200
	Fail             Code = 500
	PermissionDenied Code = 401
	SignErr          Code = 402
	ParamErr         Code = 501
	SystemErr        Code = 502
)

func (c Code) String() string {
	switch c {
	case OK:
		return "success"
	case Fail:
		return "fail"
	case ParamErr:
		return "参数错误!"
	case SignErr:
		return "签名sign错误!"
	case PermissionDenied:
		return "您还没有登录，请先登录！"
	case SystemErr:
		return "系统错误，请联系客服！"
	default:
		return "未知错误，请联系客服！"
	}
}
