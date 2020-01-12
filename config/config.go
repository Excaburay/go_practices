package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init()  {
	configFileName := "config"
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig();err != nil{
		panic(fmt.Errorf("Fatal error config file: %s \n",err))
	}
	viper.SetConfigType("yaml")

}

func GetName() string{
	return viper.GetString("name")
}

