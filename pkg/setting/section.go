package setting

import "time"

type HTTPServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HTTPLogSettingS struct {
	LogSavePath string
	LogFileExt  string
	LogFileName string
}

type TCPLogSettingS struct {
	LogSavePath string
	LogFileExt  string
	LogFileName string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	PassWord     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type TCPServerUploadSettingS struct {
	UploadSavePath     string
	UploadServerUrl    string
	UploadImageMaxSize int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type RedisSettingS struct {
	DB       int
	Password string
	Host     string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
