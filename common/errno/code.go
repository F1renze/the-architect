package errno


var (
	// sys error
	OK = newCode(0, "OK")
	SystemErr = newCode(10001, "System error")

	// 服务级错误代码, 参考 https://open.weibo.com/wiki/Error_code
	DBErr           = newCode(20607, "数据库错误，请联系系统管理员")
	RecordNotExists = newCode(20608, "找不到对应记录")

	AuthIdAlreadyUsed   = newCode(21102, "绑定失败，已被其他账户使用")
	AuthTypeAlreadyBind = newCode(21103, "不支持绑定多个相同登录方式")
	GeneratePwdErr      = newCode(21104, "不建议使用的密码")
	InvalidAuthInfo     = newCode(21105, "无效的认证信息")

	PwdNotCorrect = newCode(212301, "用户名或密码不正确")

	// mysql error number
	MySQLDupEntryErrNo = 1062
)
