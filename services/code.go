package services

/**
 * 2000+ 成功响应
 * 4000+ 请求错误
 * 5000+ 服务异常
 */

type resCode int

const (
	codeSuccess resCode = 2000 + iota
)

const (
	codeParamError resCode = 4000 + iota
	codePayloadError
	codeUsernameOrPasswordError
	codeNoRight
)

const (
	codeRefuse resCode = 5000 + iota
	codeServiceBusy
)

func (rc resCode) Msg() string {
	switch rc {
	case codeSuccess:
		return "成功访问"
	case codeParamError:
		return "参数错误"
	case codePayloadError:
		return "有效载荷部分出错"
	case codeUsernameOrPasswordError:
		return "用户名或用户密码错误"
	case codeNoRight:
		return "僭越！！！"
	case codeRefuse:
		return "拒绝服务"
	case codeServiceBusy:
		return "服务繁忙"
	default:
		return ""
	}
}
