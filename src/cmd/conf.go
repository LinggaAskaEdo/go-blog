package main

import (
	"github.com/spf13/viper"

	liblog "github.com/linggaaskaedo/go-blog/stdlib/logger"
	libredis "github.com/linggaaskaedo/go-blog/stdlib/redis"
	librouter "github.com/linggaaskaedo/go-blog/stdlib/router"
	libsql "github.com/linggaaskaedo/go-blog/stdlib/sql"
)

type Config struct {
	App    Options
	Log    liblog.Options
	Redis  libredis.Options
	SQL    map[string]libsql.Options
	Router librouter.Options
}

type Options struct {
	port int
}

func InitConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
