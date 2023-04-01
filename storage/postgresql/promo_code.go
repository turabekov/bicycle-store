package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type promoCodeRepo struct {
	db *pgxpool.Pool
}

func NewPromoCodeRepo(db *pgxpool.Pool) *promoCodeRepo {
	return &promoCodeRepo{
		db: db,
	}
}

func (r *promoCodeRepo) Create(ctx context.Context, req *models.CreatePromoCode) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO promo_code(
			promo_id,
			name,
			discount, 
			discount_type,
			order_limit_price 
		)
		VALUES (
			(
				SELECT COALESCE(MAX(promo_id), 0) + 1 FROM promo_code
			),
			$1, $2, $3, $4) RETURNING promo_id
	`
	err := r.db.QueryRow(ctx, query,
		req.Name,
		req.Discount,
		req.DiscountType,
		req.OrderLimitPrice,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *promoCodeRepo) GetByID(ctx context.Context, req *models.PromoCodePrimaryKey) (*models.PromoCode, error) {

	var (
		query     string
		promoCode models.PromoCode
	)

	query = `
		SELECT
			promo_id,
			name,
			discount, 
			discount_type,
			order_limit_price 
		FROM promo_code
		WHERE promo_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.PromoCodeId).Scan(
		&promoCode.PromoCodeId,
		&promoCode.Name,
		&promoCode.Discount,
		&promoCode.DiscountType,
		&promoCode.OrderLimitPrice,
	)
	if err != nil {
		return nil, err
	}

	return &promoCode, nil
}

func (r *promoCodeRepo) GetList(ctx context.Context, req *models.GetListPromoCodeRequest) (resp *models.GetListPromoCodeResponse, err error) {

	resp = &models.GetListPromoCodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			promo_id,
			name,
			discount, 
			discount_type,
			order_limit_price 
		FROM promo_code
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

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var promoCode models.PromoCode
		err = rows.Scan(
			&resp.Count,
			&promoCode.PromoCodeId,
			&promoCode.Name,
			&promoCode.Discount,
			&promoCode.DiscountType,
			&promoCode.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.PromoCodes = append(resp.PromoCodes, &promoCode)
	}

	return resp, nil
}

func (r *promoCodeRepo) Delete(ctx context.Context, req *models.PromoCodePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM promo_code
		WHERE promo_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.PromoCodeId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
