package conf

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings") // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf/")  // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic("Fatal error config file: " + err.Error())
	}
}
