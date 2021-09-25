package config

import (
	"github.com/spf13/viper"
	"myblog/lib/helper"
)

type StrMap map[string]interface{}

var Viper *viper.Viper

func init(){
	Viper = viper.New()
	Viper.SetConfigFile(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()

	helper.LogError(err)

	Viper.SetEnvPrefix("appenv")
	Viper.AutomaticEnv()
}

func Env(envName string, defaultData ...interface{}) interface{}{
	if !Viper.IsSet(envName){
		if len(defaultData) > 0 {
			return defaultData[0]
		}
		return nil
	}
	return Viper.Get(envName)
}


func Add(envName string, configuration map[string]interface{}){
	Viper.Set(envName, configuration)
}

