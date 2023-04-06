package postgresql

import (
	"app/config"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	categoryTestRepo  *categoryRepo
	brandTestRepo     *brandRepo
	productTestRepo   *productRepo
	stockTestRepo     *stockRepo
	customerTestRepo  *customerRepo
	storeTestRepo     *storeRepo
	staffTestRepo     *staffRepo
	orderTestRepo     *orderRepo
	promocodeTestRepo *promoCodeRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(pool)
	}

	categoryTestRepo = NewCategoryRepo(pool)
	brandTestRepo = NewBrandRepo(pool)
	productTestRepo = NewProductRepo(pool)
	stockTestRepo = NewStockRepo(pool)
	customerTestRepo = NewCustomerRepo(pool)
	storeTestRepo = NewStoreRepo(pool)
	staffTestRepo = NewStaffRepo(pool)
	orderTestRepo = NewOrderRepo(pool)
	promocodeTestRepo = NewPromoCodeRepo(pool)

	os.Exit(m.Run())
}
