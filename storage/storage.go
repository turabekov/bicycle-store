package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
	Brand() BrandRepoI
	Stock() StockRepoI
	Store() StoreRepoI
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (int, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (int, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
}

type BrandRepoI interface {
	Create(context.Context, *models.CreateBrand) (int, error)
	GetByID(context.Context, *models.BrandPrimaryKey) (*models.Brand, error)
	GetList(context.Context, *models.GetListBrandRequest) (*models.GetListBrandResponse, error)
	Update(ctx context.Context, req *models.UpdateBrand) (int64, error)
	Delete(ctx context.Context, req *models.BrandPrimaryKey) (int64, error)
}

type StockRepoI interface {
	// Create(context.Context, *models.CreateBrand) (int, error)
	GetByID(ctx context.Context, req *models.StockPrimaryKey) (*models.GetStock, error)
	GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error)
	// Update(ctx context.Context, req *models.UpdateBrand) (int64, error)
	Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error)
}

type StoreRepoI interface {
	Create(ctx context.Context, req *models.CreateStore) (int, error)
	GetByID(ctx context.Context, req *models.StorePrimaryKey) (*models.Store, error)
	GetList(ctx context.Context, req *models.GetListStoreRequest) (resp *models.GetListStoreResponse, err error)
	UpdatePut(ctx context.Context, req *models.UpdateStore) (int64, error)
	UpdatePatch(ctx context.Context, req *models.PatchRequest) (int64, error)
	Delete(ctx context.Context, req *models.StorePrimaryKey) (int64, error)
}
