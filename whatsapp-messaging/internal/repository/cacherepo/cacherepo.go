package cacherepo

import (
	"context"
	"fmt"
	"os"
	"time"
	"whatsapp-messaging/internal/logger"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var redisConnection *redis.Client

var CacheHost string = "CACHE_HOST"
var CacheSecret string = "CACHE_SECRET"
var CachePort string = "CACHE_PORT"

type DBCache[K comparable, V any] struct {
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

func (cfg redisConfig) buildConfigFromEnv() (err error) {
	cfg.Address = os.Getenv(CacheHost)
	if cfg.Address == "" {
		return fmt.Errorf("%v env must be specified: cache address", CacheHost)
	}

	cfg.Port = os.Getenv(CachePort)
	if cfg.Port == "" {
		return fmt.Errorf("%v env must be specified: cache port", CachePort)
	}

	cfg.Password = os.Getenv(CacheSecret)
	if cfg.Password == "" {
		return fmt.Errorf("%v env must be specified: cache password", CacheSecret)
	}

	return nil
}

func (cfg redisConfig) toConnOptions() *redis.Options {

	return &redis.Options{
		Network: "tcp",
		Addr:    cfg.Address + ":" + cfg.Port,
	}
}

func NewRedisClient() error {
	logger.Info(context.Background(), "connecting to redis cache...")

	cfg := &redisConfig{}
	if err := cfg.buildConfigFromEnv(); err != nil {
		return fmt.Errorf("could not load config, %w", err)
	}

	connOpt := cfg.toConnOptions()
	redisConnection = redis.NewClient(connOpt)

	return nil
}

func MakeCache[K comparable, V any](ExpTime time.Duration, loader Loader[K, V]) *DBCache[K, V] {
	return &DBCache[K, V]{
		rCache: cache.New(&cache.Options{
			Redis: redisConnection,
		}),
		Loader: loader,
		TTL:    ExpTime,
	}
}

func (c DBCache[K, V]) Read(ctx context.Context, key K) (value V, err error) {
	if c.Loader == nil {
		logger.Panic(ctx, "cache is not initialized")
	}

	if err = c.rCache.Get(ctx, fmt.Sprint(key), value); err == nil {
		return value, nil
	}
	logger.Error(ctx, err)

	value, err = c.Loader(ctx, key)
	if err != nil {
		var v V
		return v, err
	}

	return value, nil
}
