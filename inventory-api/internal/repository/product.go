package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/bobbybof/inventory-api/internal/helper"
)

type GetAllProductsParam struct {
	Limit  int64  `form:"limit"`
	Offset int64  `form:"offset"`
	Search string `form:"search" default:"''"`
}

func (q *Queries) GetAllProducts(ctx context.Context, params GetAllProductsParam) ([]Product, int64, error) {
	query := helper.QueryBuilder.Select("*").From("products").Limit(uint64(params.Limit)).Offset(uint64(params.Offset))

	if params.Search != "" {
		query = query.Where(sq.Expr("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search)))
	}

	sql, args, err := query.ToSql()

	fmt.Println(sql, args)

	if err != nil {
		return nil, 0, err
	}

	count, err := q.CountProduct(ctx)
	if err != nil {
		return nil, 0, err
	}

	rows, err := q.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, count, nil

}
