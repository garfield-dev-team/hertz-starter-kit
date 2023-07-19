package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func SetupConfig() {
	var err error
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
	if err = vp.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
}
