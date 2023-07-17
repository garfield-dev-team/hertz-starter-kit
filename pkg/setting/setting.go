package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string        `json:"run_mode" mapstructure:"run_mode"`
	HttpPort     string        `json:"http_port" mapstructure:"http_port"`
	ReadTimeout  time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
}

type AppSettingS struct {
	DefaultPageSize int    `json:"default_page_size" mapstructure:"default_page_size"`
	MaxPageSize     int    `json:"max_page_size" mapstructure:"max_page_size"`
	LogSavePath     string `json:"log_save_path" mapstructure:"log_save_path"`
	LogFileName     string `json:"log_file_name" mapstructure:"log_file_name"`
	LogFileExt      string `json:"log_file_ext" mapstructure:"log_file_ext"`
}

type DatabaseSettingS struct {
	Type         string `json:"type"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Name         string `json:"name"`
	TablePrefix  string `json:"table_prefix" mapstructure:"table_prefix"`
	Charset      string `json:"charset"`
	ParseTime    bool   `json:"parse_time" mapstructure:"parse_time"`
	MaxIdleConns int    `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"`
}

var (
	ServerSetting   *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
)
