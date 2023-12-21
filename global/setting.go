package global

import "github.com/SLOWDOWNO/blog-service-go/pkg/setting"

// global variable
// 将配置文件和程序关联起来
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
