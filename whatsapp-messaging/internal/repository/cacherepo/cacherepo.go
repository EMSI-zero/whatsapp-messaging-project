package cacherepo

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"whatapp-messaging/internal/logger"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var redisConnection *redis.Client

type dbCache[K comparable, V any] struct {
	rCache *cache.Cache
	Loader Loader[K, V]
	TTL    time.Duration
}

type Loader[K comparable, V any] func(ctx context.Context, key K) (value V, err error)

type redisConfig struct {
	Address  string
	Port     string
	Password string
}

func (cfg redisConfig) buildConfigFromEnv() error {
	return nil
}

func (cfg redisConfig) toConnOptions() *redis.Options {

	return &redis.Options{
		Network: "tcp",
		Addr:    cfg.Address + ":" + cfg.Port,
	}
}

func NewRedisClient() error {
	logger.Info("connecting to redis cache...")

	cfg := &redisConfig{}
	if err := cfg.buildConfigFromEnv(); err != nil {
		return fmt.Errorf("could not load config, %w", err)
	}

	connOpt := cfg.toConnOptions()
	redisConnection = redis.NewClient(connOpt)

	return nil
}

func (c dbCache[K, V]) MakeCache(ExpTime time.Duration, loader Loader[K, V]) *dbCache[K, V] {
	return &dbCache[K, V]{
		rCache: cache.New(&cache.Options{
			Redis: redisConnection,
		}),
		Loader: loader,
		TTL:    ExpTime,
	}
}

func (c dbCache[K, V]) Read(ctx context.Context, key K) (value V,err error) {
	if c.Loader == nil {
		logger.Panic("cache is not initialized")
	}

	if err = c.rCache.Get(ctx, fmt.Sprint(key) , value); err == nil{
		return value, nil
	}
	logger.Error(err)

	value, err = c.Loader(ctx, key)
	if err != nil {
		var v V
		return v, err
	}

	return value, nil
}
