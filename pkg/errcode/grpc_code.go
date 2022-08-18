package errcode

var (
	ErrorPassword = NewError(30010001, "账号或密码错误")
	ErrorAuth     = NewError(30010002, "验证失败")
	ErrorOther    = NewError(30010003, "其他错误")
)
