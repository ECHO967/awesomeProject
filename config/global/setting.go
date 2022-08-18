package global

import (
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/setting"
)

var (
	HTTPServerSetting *setting.HTTPServerSettingS

	DataBaseSetting *setting.DatabaseSettingS

	RedisSetting *setting.RedisSettingS

	HTTPLogSetting *setting.HTTPLogSettingS

	TCPLogSetting *setting.TCPLogSettingS

	UploadSetting *setting.TCPServerUploadSettingS

	JWTSetting *setting.JWTSettingS

	//ServerSetting *setting.ServerSetting

	HTTPLogger *logger.Logger
	TCPLogger  *logger.Logger
)
