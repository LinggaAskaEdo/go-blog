package main

import (
	"github.com/spf13/viper"

	liblog "github.com/linggaaskaedo/go-blog/stdlib/logger"
	libmiddleware "github.com/linggaaskaedo/go-blog/stdlib/middleware"
	libmigrate "github.com/linggaaskaedo/go-blog/stdlib/migrate"
	libmux "github.com/linggaaskaedo/go-blog/stdlib/mux"
	libparser "github.com/linggaaskaedo/go-blog/stdlib/parser"
	libredis "github.com/linggaaskaedo/go-blog/stdlib/redis"
	libhttpserver "github.com/linggaaskaedo/go-blog/stdlib/server"
	libsql "github.com/linggaaskaedo/go-blog/stdlib/sql"
)

type Config struct {
	App        Options
	Log        liblog.Options
	Redis      libredis.Options
	SQL        map[string]libsql.Options
	Middleware libmiddleware.Options
	Mux        libmux.Options
	Server     libhttpserver.Options
	Migrate    libmigrate.Options
	Parser     libparser.Options
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
