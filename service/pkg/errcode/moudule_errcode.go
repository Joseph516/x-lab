package errcode

// 功能服务错误码

var (
	// UploadFileFailed 文件上传
	UploadFileFailed = NewError(2001, "上传文件失败")

	PasswordNotSame = NewError(2002, "两次密码不一致")
	RegisterFailed  = NewError(2003, "注册失败")
	LoginInFailed   = NewError(2004, "登陆失败，用户或密码错误")
)
