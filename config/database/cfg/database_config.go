package cfg

import "github.com/spf13/viper"

var cfg *DbConfig

type DbConfig struct {
	Host     string
	User     string
	PassWord string
	DataBase string
	Port     string
}

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	cfg = &DbConfig{
		Host:     viper.GetString("db.host"),
		User:     viper.GetString("db.user"),
		PassWord: viper.GetString("db.password"),
		DataBase: viper.GetString("db.database"),
		Port:     viper.GetString("db.port"),
	}
}

func GetConfigDb() DbConfig {
	Load()
	return *cfg
}
