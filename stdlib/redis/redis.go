package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type Options struct {
	Enabled         bool
	Address         []string
	Password        string
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PoolSize        int
	MinIdleConns    int
	MaxIdleConns    int
	MaxActiveConns  int
	PoolTimeout     time.Duration
	MaxRedirects    int
	ReadOnly        bool
	RouteByLatency  bool
	RouteRandomly   bool
}

func Init(log zerolog.Logger, opt Options) *redis.Client {
	var redisClient *redis.Client

	if !opt.Enabled {
		return nil
	}

	once.Do(func() {
		univOptions := &redis.UniversalOptions{
			Addrs:           opt.Address,
			Password:        opt.Password,
			MaxRetries:      opt.MaxRetries,
			MinRetryBackoff: opt.MinRetryBackoff,
			MaxRetryBackoff: opt.MaxRetryBackoff,
			DialTimeout:     opt.DialTimeout,
			ReadTimeout:     opt.ReadTimeout,
			WriteTimeout:    opt.WriteTimeout,
			PoolSize:        opt.PoolSize,
			MinIdleConns:    opt.MinIdleConns,
			MaxIdleConns:    opt.MaxIdleConns,
			MaxActiveConns:  opt.MaxActiveConns,
			PoolTimeout:     opt.PoolTimeout,
			MaxRedirects:    opt.MaxRedirects,
			ReadOnly:        opt.ReadOnly,
			RouteByLatency:  opt.RouteByLatency,
			RouteRandomly:   opt.RouteRandomly,
		}

		univClient := redis.NewUniversalClient(univOptions)
		redisClient := univClient.(*redis.Client)

		ping, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Panic().Err(err).Str("redis_status", "FAILED").Send()
		}

		log.Debug().Str("redis_status", ping).Send()
	})

	return redisClient
}
