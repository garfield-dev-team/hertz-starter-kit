package setting

import (
	"time"

	"github.com/spf13/viper"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

var (
	ServerSetting   *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
)

func SetupSetting() (err error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		return
	}
	if err = vp.UnmarshalKey("Server", &ServerSetting); err != nil {
		return
	}
	if err = vp.UnmarshalKey("App", &AppSetting); err != nil {
		return
	}
	if err = vp.UnmarshalKey("Database", &DatabaseSetting); err != nil {
		return
	}
	return
}
