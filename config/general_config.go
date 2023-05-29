package config

import "github.com/spf13/viper"

var cfgGeneral *GeneralConfig

type GeneralConfig struct {
	PortApi string
}

func LoadCfgGeneral() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	cfgGeneral = &GeneralConfig{
		PortApi: viper.GetString("api.port"),
	}
}

func GetGeneralConfig() GeneralConfig {
	return *cfgGeneral
}
