package common

// 全局返回码
// API https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Global_Return_Code.html
const (
	ErrCodeCommonSystemIsBusy      = -1    // 系统繁忙，此时请开发者稍候再试
	ErrCodeCommonSuccess           = 0     // 请求成功
	ErrCodeCommonAPIfReqOutOfLimit = 45009 // api频率超出限制
	ErrCodeCommonInvalidCredential = 40001 // 无效的凭证类型
	ErrCodeCommonIllegalCredential = 40002 // 不合法的凭证类型
)
