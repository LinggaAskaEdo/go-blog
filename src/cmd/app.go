package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	domain "github.com/linggaaskaedo/go-blog/src/business/domain"
	usecase "github.com/linggaaskaedo/go-blog/src/business/usecase"
	resthandler "github.com/linggaaskaedo/go-blog/src/handler/rest"
	libgrace "github.com/linggaaskaedo/go-blog/stdlib/grace"
	liblog "github.com/linggaaskaedo/go-blog/stdlib/logger"
	libmux "github.com/linggaaskaedo/go-blog/stdlib/mux"
	libredis "github.com/linggaaskaedo/go-blog/stdlib/redis"
	libhttpserver "github.com/linggaaskaedo/go-blog/stdlib/server"
	libsql "github.com/linggaaskaedo/go-blog/stdlib/sql"
)

var (
	confPath  string
	minJitter int
	maxJitter int

	dom *domain.Domain
	uc  *usecase.Usecase

	redisClient0 *redis.Client
	sqlClient0   *sqlx.DB
	sqlClient1   *sqlx.DB
	httpMux      *mux.Router
	httpServer   *http.Server
	app          libgrace.App
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

	// HTTP Mux Initialization
	httpMux = libmux.Init(logger, conf.Mux)

	// Domain Initialization
	dom = domain.Init(logger, redisClient0, sqlClient0, sqlClient1)

	// Usecase Initialization
	uc = usecase.Init(logger, redisClient0, sqlClient0, sqlClient1, dom)

	// REST Handler Initialization
	_ = resthandler.Init(logger, httpMux, uc)

	// HTTP Server Initialization
	httpServer = libhttpserver.Init(logger, conf.Server, httpMux)

	// App Initialization
	app = libgrace.Init(logger, httpServer)
}

func main() {
	defer func() {
		if redisClient0 != nil {
			redisClient0.Close()
		}

		if sqlClient0 != nil {
			sqlClient0.Close()
		}

		if sqlClient1 != nil {
			sqlClient1.Close()
		}
	}()

	app.Serve()
}
