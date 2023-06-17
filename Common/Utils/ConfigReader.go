package Utils

import (
	"github.com/spf13/viper"
	"log"
)

func ReadConfigVal(key string) any {
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yml")
	// 添加配置文件的路径，指定在 Config 目录下寻找
	viper.AddConfigPath("./Config")

	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fail to read config file: %e", err)
	}
	return viper.Get(key)
}
