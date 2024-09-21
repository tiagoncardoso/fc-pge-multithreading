package config

import "github.com/spf13/viper"

type conf struct {
	ApiRequest1URL string `mapstructure:"API_REQUEST_1_URL"`
	ApiRequest2URL string `mapstructure:"API_REQUEST_2_URL"`
	RequestTimeout int    `mapstructure:"REQUEST_TIMEOUT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
