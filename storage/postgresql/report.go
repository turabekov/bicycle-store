package postgresql

import (
	"app/api/models"
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type reportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *reportRepo {
	return &reportRepo{
		db: db,
	}
}

// task1
func (r *reportRepo) ExchangeStoreProduct(ctx context.Context, req *models.ExchangeProduct) (fromId int, toId int, err error) {

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	_, err = tx.Exec(ctx, `UPDATE stocks SET quantity = quantity - $1 WHERE store_id = $2 AND product_id = $3`, req.Quantity, req.FromStoreId, req.ProductId)
	if err != nil {
		tx.Rollback(ctx)
		return 0, 0, err
	}
	_, err = tx.Exec(ctx, `UPDATE stocks SET quantity = quantity + $1 WHERE store_id = $2 AND product_id = $3`, req.Quantity, req.ToStoreId, req.ProductId)
	if err != nil {
		tx.Rollback(ctx)
		return 0, 0, err
	}

	return req.FromStoreId, req.ToStoreId, nil
}
