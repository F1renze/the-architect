package errno

var (
	// sys error
	OK               = newCode(0, "OK")
	SystemErr        = newCode(10001, "System error")
	RemoteServiceErr = newCode(10003, "Remote service error")

	// 服务级错误代码, 参考 https://open.weibo.com/wiki/Error_code
	DBErr           = newCode(20607, "数据库错误，请联系系统管理员")
	RecordNotExists = newCode(20608, "找不到对应记录")

	RedisOperationErr = newCode(20701, "Redis 操作失败，请联系系统管理员")

	// 认证相关
	AuthIdAlreadyUsed   = newCode(21102, "绑定失败，已被其他账户使用")
	AuthTypeAlreadyBind = newCode(21103, "不支持绑定多个相同登录方式")
	GeneratePwdErr      = newCode(21104, "不建议使用的密码")
	InvalidAuthInfo     = newCode(21105, "无效的认证信息")
	MobileAlreadyUsed   = newCode(21106, "该手机号已被使用")

	// user 相关（含登录注册
	InvalidUserName     = newCode(21201, "用户名不合法")
	UserNameAlreadyUsed = newCode(21202, "用户名已被使用")
	UidIsNull           = newCode(21203, "uid 为空")
	InvalidEmail        = newCode(21204, "无效的邮箱地址")

	PwdNotCorrect = newCode(212301, "用户名或密码不正确")

	// token 相关
	SignedTokenFailed       = newCode(213401, "Token 签名失败")
	InvalidToken            = newCode(213402, "Token 不合法")
	SignatureMethodRejected = newCode(213403, "签名算法不支持")
	TokenExpired            = newCode(213404, "Token 已失效")

	// casbin
	EnforcerNotFound = newCode(213501, "Enforcer Not Found")
	AdapterNotFound = newCode(213502, "Adapter Not Found")
	InvalidModel = newCode(213503, "Invalid casbin model")
	UnSupportAdapterDriver = newCode(213504, "driver unsupported")
	CreateAdapterErr = newCode(213505, "create adapter error")

	// mysql error number
	MySQLDupEntryErrNo = 1062
)
