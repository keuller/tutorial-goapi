package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Get() *viper.Viper {
	cfg := viper.New()
	cfg.AddConfigPath(".")
	// cfg.AddConfigPath("$HOME/workspace/estudos/tutorial-goapi")
	cfg.SetConfigName("env")
	cfg.SetConfigType("yaml")

	if err := cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error on reading configuration file %s\n", err))
	}

	return cfg
}
