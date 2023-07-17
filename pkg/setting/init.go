package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

func SetupSetting() {
	var err error
	defer func() {
		if err != nil {
			panic(fmt.Errorf("fatal error configs file: %w", err))
		}
	}()

	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		return
	}
	if err = vp.UnmarshalKey("server", &ServerSetting); err != nil {
		return
	}
	if err = vp.UnmarshalKey("app", &AppSetting); err != nil {
		return
	}
	if err = vp.UnmarshalKey("database", &DatabaseSetting); err != nil {
		return
	}
}
