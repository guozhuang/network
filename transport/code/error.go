package code

//错误码统一标注结构

const HttpOk = 200
const HttpUserError = 403
const HttpLoginExpire = 406
const HttpLoginFail = 407

const MoaOk = 0

var ErrMap = map[int]string{
	HttpOk:          "",
	HttpUserError:   "该请求错误",
	HttpLoginExpire: "登陆已过期",
	HttpLoginFail:   "登陆验签失败",
}
