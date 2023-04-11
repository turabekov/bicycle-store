package redis

import (
	"app/config"
	"app/storage"

	"github.com/go-redis/redis"
)

type CacheStore struct {
	redisDB          *redis.Client
	productCacheRepo storage.ProductCacheRepoI
}

func NewConnectRedis(cfg *config.Config) (storage.StorageCacheI, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     cfg.RedisHost + cfg.RedisPort,
			Password: "",
			DB:       0,
		},
	)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &CacheStore{
		redisDB:          client,
		productCacheRepo: NewProductCacheRepo(client),
	}, nil
}

func (c *CacheStore) CloseDB() {
	c.redisDB.Close()
}

func (c *CacheStore) Product() storage.ProductCacheRepoI {

	if c.productCacheRepo == nil {
		c.productCacheRepo = NewProductCacheRepo(c.redisDB)
	}

	return c.productCacheRepo
}
