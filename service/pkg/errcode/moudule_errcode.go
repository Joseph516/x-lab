package errcode

// 功能服务错误码

var (
	// UploadFileFailed 文件上传
	UploadFileFailed = NewError(20001, "上传文件失败")
)
