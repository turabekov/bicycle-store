package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

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

// task2
func (r *reportRepo) StaffSaleReport(ctx context.Context, req *models.GetListEmployeeReportRequest) (resp *models.GetListEmployeeReportResponse, err error) {
	resp = &models.GetListEmployeeReportResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT
		(sta.first_name || ' ' || sta.last_name) AS employee,
    	sto.store_name,
    	c.category_name,
    	p.product_name,
    	(oi.quantity) AS total_amount,
    	(oi.list_price) * (oi.quantity)  AS total_price,
    	CAST(o.order_date::timestamp AS VARCHAR)
	FROM orders AS o
	JOIN staffs AS sta ON sta.staff_id = o.staff_id 
	JOIN stores AS sto ON sto.store_id = o.store_id 
	JOIN order_items AS oi ON oi.order_id = o.order_id  
	JOIN products AS p ON oi.product_id = p.product_id
	JOIN categories AS c ON c.category_id = p.category_id

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

	query += filter + " ORDER BY o.order_date " + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		var employee models.EmployeeReport

		err = rows.Scan(
			&employee.EmployeeFullName,
			&employee.StoreName,
			&employee.CategoryName,
			&employee.ProductName,
			&employee.Quantity,
			&employee.TotalPrice,
			&employee.Date,
		)
		if err != nil {
			return nil, err
		}

		resp.EmployeeReports = append(resp.EmployeeReports, &employee)

		counter++
	}
	resp.Count = counter

	return resp, nil
}
