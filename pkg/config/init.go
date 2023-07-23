package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func SetupConfig() {
	var err error
	var configPath string

	env := os.Getenv("GO_ENV")
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}
	if env == "test" {
		configPath = filepath.Join(wd, "../../../configs/")
	} else {
		configPath = filepath.Join(wd, "configs/")
	}

	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(configPath)
	vp.SetConfigType("yaml")
	if err = vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
	if err = vp.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}
}
