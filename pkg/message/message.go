package message

const (
	CloseFileError = "关闭文件错误"
)

const (
	LogConfigError      = "配置日志错误"
	DatabaseConfigError = "配置数据库错误"
	AppConfigError      = "配置应用错误"

	AppConfigOk = "配置成功"
)

const (
	DatabaseConnectError = "数据库连接出错"
	GetConnectionError   = "获取当前数据库连接出错"
	ORMConfigError       = "配置ORM错误"
	DatabaseMigrateError = "数据库迁移出错"
)

const (
	HTTPServerStartError = "HTTP服务启动出错"
	HTTPServerStopError  = "HTTP服务停止出错"
)

const (
	TransGetTranslatorError = "根据语言选择翻译器出错"
	TransRegisterError      = "国际化出错"
)

const (
	MinioConfigError           = "Minio配置错误"
	MinioCheckBucketExistError = "Minio检查bucket是否存在出错"
	MinioMakeBucketError       = "Minio创建bucket出错"
	MinioReadFileError         = "Minio读取文件错误"
	MinioPutObjectError        = "Minio创建对象失败"
	MinioSetPolicyError        = "Minio创建Policy失败"

	MinioMakeBucketOk = "Minio创建bucket成功"
)

const (
	ConsulConfigError = "Consul配置错误"
)
