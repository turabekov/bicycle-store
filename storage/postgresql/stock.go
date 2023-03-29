package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type stockRepo struct {
	db *pgxpool.Pool
}

func NewStockRepo(db *pgxpool.Pool) *stockRepo {
	return &stockRepo{
		db: db,
	}
}

func (r *stockRepo) Create(ctx context.Context, req *models.CreateStock) (int, int, error) {
	var (
		query     string
		storeId   int
		productId int
	)

	query = `
		INSERT INTO stocks(
			store_id,
			product_id,
			quantity
		)
		VALUES ($1, $2, $3) RETURNING store_id, product_id
	`
	err := r.db.QueryRow(ctx, query,
		req.StoreId,
		req.ProductId,
		req.Quantity,
	).Scan(&storeId, productId)
	if err != nil {
		return 0, 0, err
	}

	return storeId, productId, nil
}

func (r *stockRepo) GetByID(ctx context.Context, req *models.StockPrimaryKey) (*models.GetStock, error) {

	var (
		query      string
		stock      models.GetStock
		productIds []sql.NullInt64
		amounts    []sql.NullInt64
	)

	if req.ProductId > 0 {
		query = `
			SELECT
				store_id,
				product_id,
				quantity
			FROM stocks
			WHERE store_id = $1 AND product_id = $2
		`
		fmt.Println(query)
		stock.Products = append(stock.Products, &models.ProductData{})
		err := r.db.QueryRow(ctx, query, req.StoreId, req.ProductId).Scan(
			&stock.StoreId,
			&stock.Products[0].ProductId,
			&stock.Products[0].Quantity,
		)
		fmt.Println("Hello")
		if err != nil {
			return nil, err
		}
		fmt.Println(stock)
		return &stock, nil
	}

	query = `
		SELECT
			store_id,
			ARRAY_AGG(product_id),
			ARRAY_AGG(quantity)
		FROM stocks
		WHERE store_id = $1
		GROUP BY store_id
	`
	err := r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&stock.StoreId,
		pq.Array(&productIds),
		pq.Array(&amounts),
	)
	if err != nil {
		return nil, err
	}

	for i, id := range productIds {
		data := models.ProductData{
			ProductId: int(id.Int64),
			Quantity:  int(amounts[i].Int64),
		}
		stock.Products = append(stock.Products, &data)
	}

	fmt.Println(stock)
	return &stock, nil
}

func (r *stockRepo) GetList(ctx context.Context, req *models.GetListStockRequest) (resp *models.GetListStockResponse, err error) {

	resp = &models.GetListStockResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			store_id,
			ARRAY_AGG(product_id),
			ARRAY_AGG(quantity)
		FROM stocks
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + " GROUP BY store_id " + offset + limit
	fmt.Println(query)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			stock      models.GetStock
			productIds []sql.NullInt64
			amounts    []sql.NullInt64
		)

		err = rows.Scan(
			&resp.Count,
			&stock.StoreId,
			pq.Array(&productIds),
			pq.Array(&amounts),
		)
		if err != nil {
			return nil, err
		}

		for i, id := range productIds {
			data := models.ProductData{
				ProductId: int(id.Int64),
				Quantity:  int(amounts[i].Int64),
			}
			stock.Products = append(stock.Products, &data)
		}

		resp.Stocks = append(resp.Stocks, &stock)
	}

	return resp, nil
}

func (r *stockRepo) Update(ctx context.Context, req *models.UpdateStock) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		stocks
		SET
			quantity = :quantity
		WHERE store_id = :store_id AND product_id = :product_id
	`

	params = map[string]interface{}{
		"store_id":   req.StoreId,
		"product_id": req.ProductId,
		"quantity":   req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *stockRepo) Delete(ctx context.Context, req *models.StockPrimaryKey) (int64, error) {

	var (
		storeId   string
		productId string
	)
	if req.StoreId > 0 {
		storeId = fmt.Sprintf(" store_id = %d ", req.StoreId)
	}

	query := `
		DELETE
		FROM stocks
		WHERE 
	` + storeId

	if req.ProductId > 0 {
		productId = fmt.Sprintf(" product_id = %d ", req.ProductId)
		query += "AND" + productId
	}

	result, err := r.db.Exec(ctx, query, req.StoreId, req.ProductId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
