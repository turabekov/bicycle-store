package storage

import "app/api/models"

type StorageCacheI interface {
	CloseDB()
	Product() ProductCacheRepoI
}

type ProductCacheRepoI interface {
	Create(*models.GetListProductResponse) error
	GetAll() (*models.GetListProductResponse, error)
	Exists(limit, offset int) (bool, error)
	Delete() error
}
