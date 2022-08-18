package errcode

var (
	ErrorGetUserFail    = NewError(20010001, "获取用户信息失败")
	ErrorChangeNickFail = NewError(20010002, "更改nickname失败")
	ErrorChangeProfFail = NewError(20010003, "上传profile失败")
)
