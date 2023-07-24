package configs

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var (
	config Config
	once sync.Once
	path = "../"
)

func Get() *Config {
	once.Do(func() {
		viper.AddConfigPath(path)

		viper.SetConfigName("app")
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	
		err = viper.Unmarshal(&config)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Configuration: ", config )
	})

	return &config
}

// func LoadConfig(path string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }
