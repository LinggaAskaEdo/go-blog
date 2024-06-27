package main

import (
	"context"
	"flag"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	domain "github.com/linggaaskaedo/go-blog/src/business/domain"
	usecase "github.com/linggaaskaedo/go-blog/src/business/usecase"
	liblog "github.com/linggaaskaedo/go-blog/stdlib/logger"
	libredis "github.com/linggaaskaedo/go-blog/stdlib/redis"
	librouter "github.com/linggaaskaedo/go-blog/stdlib/router"
	libsql "github.com/linggaaskaedo/go-blog/stdlib/sql"
)

var (
	confPath  string
	minJitter int
	maxJitter int

	redisClient0 *redis.Client
	sqlClient0   *sqlx.DB
	sqlClient1   *sqlx.DB
	httpRouter   *mux.Router

	dom *domain.Domain
	uc  *usecase.Usecase
)

func init() {
	// Flag Settings Initialization
	flag.StringVar(&confPath, "staticConfPath", "./etc/conf", "config path")
	flag.IntVar(&minJitter, "minSleep", DefaultMinJitter, "min. sleep duration during app initialization")
	flag.IntVar(&maxJitter, "maxSleep", DefaultMaxJitter, "max. sleep duration during app initialization")
	flag.Parse()

	// Add Sleep with Jitter to drag the the initialization time among instances
	sleepWithJitter(minJitter, maxJitter)

	// Config Initialization
	conf, err := InitConfig(confPath)
	if err != nil {
		panic(err)
	}

	// Logger Initialization
	logger := liblog.Init(conf.Log)

	// Redis Initialization
	redisClient0 = libredis.Init(logger, conf.Redis)

	// SQL Initialization
	sqlClient0 = libsql.Init(logger, conf.SQL["sql-0"])
	sqlClient1 = libsql.Init(logger, conf.SQL["sql-1"])

	// Router Initialization
	httpRouter = librouter.Init(logger, conf.Router)

	// Domain Initialization
	dom = domain.Init(logger, redisClient0, sqlClient0, sqlClient1)

	// Usecase Initialization
	uc = usecase.Init(logger, redisClient0, sqlClient0, sqlClient1, dom)
}

func main() {
	defer func() {
		if redisClient0 != nil {
			redisClient0.Conn().Quit(context.Background())
		}

		if sqlClient0 != nil {
			sqlClient0.Close()
		}

		if sqlClient1 != nil {
			sqlClient1.Close()
		}
	}()
}
