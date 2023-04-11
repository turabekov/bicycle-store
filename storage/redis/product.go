package redis

import (
	"app/api/models"
	"encoding/json"

	"github.com/go-redis/redis"
)

type productCacheRepo struct {
	cache *redis.Client
}

func NewProductCacheRepo(redisDB *redis.Client) *productCacheRepo {
	return &productCacheRepo{
		cache: redisDB,
	}
}

func (c *productCacheRepo) Create(req *models.GetListProductResponse) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.cache.Set("products", body, 0).Err()
	if err != nil {
		return err
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}
	err = c.cache.Set("product_limit", req.Limit, 0).Err()
	if err != nil {
		return err
	}

	if req.Offset <= 0 {
		req.Offset = 0
	}

	err = c.cache.Set("product_offset", req.Offset, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *productCacheRepo) GetAll() (*models.GetListProductResponse, error) {

	resp := models.GetListProductResponse{}

	productData, err := c.cache.Get("products").Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(productData), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *productCacheRepo) Exists(limit, offset int) (bool, error) {

	exists, err := c.cache.Exists("products").Result()
	if err != nil {
		return false, err
	}

	if exists > 0 {
		resp := models.GetListProductResponse{}

		productData, err := c.cache.Get("products").Result()
		if err != nil {
			return false, err
		}

		err = json.Unmarshal([]byte(productData), &resp)
		if err != nil {
			return false, err
		}

		if resp.Limit != limit {
			return false, nil
		}

		if resp.Offset != offset {
			return false, nil
		}

	}

	if exists <= 0 {
		return false, nil
	}

	return true, nil
}

func (c *productCacheRepo) Delete() error {

	err := c.cache.Del("products").Err()
	if err != nil {
		return err
	}

	return nil
}
