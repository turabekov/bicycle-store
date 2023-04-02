package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
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
	).Scan(&storeId, &productId)
	if err != nil {
		return 0, 0, err
	}

	return storeId, productId, nil
}

func (r *stockRepo) GetByIdProductStock(ctx context.Context, storeId int, productId int) (resp *models.Stock, err error) {
	resp = &models.Stock{}
	resp.StoreData = &models.Store{}
	resp.ProductData = &models.Product{}

	query := `
		SELECT 
			s.store_id,

			st.store_id, 
			st.store_name,
			COALESCE(st.phone, ''),
			COALESCE(st.email, ''),
			COALESCE(st.street, ''),
			COALESCE(st.city, ''),
			COALESCE(st.state, ''),
			COALESCE(st.zip_code, ''),

			s.product_id,

			p.product_id, 
			p.product_name, 
			p.brand_id,
			p.category_id,
			p.model_year,
			p.list_price,
			
			s.quantity
		FROM stocks AS s
		JOIN stores AS st ON st.store_id = s.store_id
		JOIN products AS p ON p.product_id = s.product_id
		WHERE s.store_id = $1 AND s.product_id = $2
	`

	err = r.db.QueryRow(ctx, query, storeId, productId).Scan(
		&resp.StoreId,
		&resp.StoreData.StoreId,
		&resp.StoreData.StoreName,
		&resp.StoreData.Phone,
		&resp.StoreData.Email,
		&resp.StoreData.Street,
		&resp.StoreData.City,
		&resp.StoreData.State,
		&resp.StoreData.ZipCode,
		&resp.ProductId,
		&resp.ProductData.ProductId,
		&resp.ProductData.ProductName,
		&resp.ProductData.BrandId,
		&resp.ProductData.CategoryId,
		&resp.ProductData.ModelYear,
		&resp.ProductData.ListPrice,
		&resp.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (r *stockRepo) GetByID(ctx context.Context, req *models.StockPrimaryKey) (resp *models.GetStock, err error) {
	resp = &models.GetStock{}

	var (
		quantity sql.NullInt64
		storeId  sql.NullInt64
		products pgtype.JSONB
	)

	query := `
		SELECT
		s.store_id,
		SUM(s.quantity),
		JSONB_AGG (
	    		JSONB_BUILD_OBJECT (
	    			'product_id', p.product_id,
				'product_name', p.product_name,
				'brand_id', p.brand_id,
				'category_id', p.category_id,
	            'category_data',    JSONB_BUILD_OBJECT(
	                'category_id', c.category_id,
	                'category_name', c.category_name
	            ),
				'model_year', p.model_year,
				'list_price', p.list_price,
				'quantity', s.quantity
			)
		) AS product_data
		FROM stocks AS s 
		LEFT JOIN products AS p ON p.product_id = s.product_id
		LEFT JOIN categories AS c ON c.category_id = p.category_id
		WHERE s.store_id = $1
		GROUP BY s.store_id
	`
	err = r.db.QueryRow(ctx, query, req.StoreId).Scan(
		&storeId,
		&quantity,
		&products,
	)
	if err != nil {
		return nil, err
	}

	resp.StoreId = int(storeId.Int64)
	resp.Quantity = int(quantity.Int64)

	products.AssignTo(&resp.Products)

	return resp, nil
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
			s.store_id,
			SUM(s.quantity),
			JSONB_AGG (
				JSONB_BUILD_OBJECT (
					'product_id', p.product_id,
					'product_name', p.product_name,
					'brand_id', p.brand_id,
					'category_id', p.category_id,
					'model_year', p.model_year,
					'list_price', p.list_price,
					'quantity', s.quantity
				)
			) AS product_data
		FROM stocks AS s 
		LEFT JOIN products AS p ON p.product_id = s.product_id

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

	query += filter + " GROUP BY s.store_id " + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			stock    models.GetStock
			quantity sql.NullInt64
			storeId  sql.NullInt64
			products pgtype.JSONB
		)

		err = rows.Scan(
			&resp.Count,
			&storeId,
			&quantity,
			&products,
		)
		if err != nil {
			return nil, err
		}

		stock.StoreId = int(storeId.Int64)
		stock.Quantity = int(quantity.Int64)

		products.AssignTo(&stock.Products)

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
		storeId string
	)
	if req.StoreId > 0 {
		storeId = fmt.Sprintf(" store_id = %d ", req.StoreId)
	}

	query := `
		DELETE
		FROM stocks
		WHERE 
	` + storeId

	result, err := r.db.Exec(ctx, query, req.StoreId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
