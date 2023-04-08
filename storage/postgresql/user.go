package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		query string
		id    string
	)
	id = uuid.NewString()

	query = `
		INSERT INTO users(
			id, 
			name,
			login,
			password,
			updated_at 
		)
		VALUES ( $1, $2, $3, $4, now())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Login,
		req.Password,
	)
	if err != nil {
		fmt.Println(err.Error())
	
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		user  models.User
	)

	if len(req.Login) > 0 {
		err := r.db.QueryRow(ctx, "SELECT id FROM users WHERE login = $1", req.Login).Scan(&req.Id)
		if err != nil {
			return nil, err
		}
	}

	query = `
		SELECT
			id, 
			name,
			login,
			password,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&user.Id,
		&user.Name,
		&user.Login,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			name,
			login,
			password,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM users
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
		var user models.User
		err = rows.Scan(
			&resp.Count,
			&user.Id,
			&user.Name,
			&user.Login,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		users
		SET
			id = :id,
			name = :name,
			login = :login,
			password = :password,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":       req.Id,
		"name":     req.Name,
		"login":    req.Login,
		"password": req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
