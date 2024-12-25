// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO catalog.category(category_id, category_name, category_description)
VALUES ($1, $2, $3)
RETURNING category_id, category_name, category_description
`

type CreateCategoryParams struct {
	CategoryID          pgtype.UUID `json:"category_id"`
	CategoryName        pgtype.Text `json:"category_name"`
	CategoryDescription pgtype.Text `json:"category_description"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (CatalogCategory, error) {
	row := q.db.QueryRow(ctx, createCategory, arg.CategoryID, arg.CategoryName, arg.CategoryDescription)
	var i CatalogCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription)
	return i, err
}

const deleteCategoryById = `-- name: DeleteCategoryById :exec
DELETE
FROM catalog.category
WHERE category_id = $1
`

func (q *Queries) DeleteCategoryById(ctx context.Context, categoryID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCategoryById, categoryID)
	return err
}

const findAllCategories = `-- name: FindAllCategories :many
SELECT category_id, category_name, category_description
FROM catalog.category
LIMIT $1 OFFSET $2
`

type FindAllCategoriesParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) FindAllCategories(ctx context.Context, arg FindAllCategoriesParams) ([]CatalogCategory, error) {
	rows, err := q.db.Query(ctx, findAllCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CatalogCategory{}
	for rows.Next() {
		var i CatalogCategory
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllChildrenCategories = `-- name: FindAllChildrenCategories :many
SELECT category_id, category_name, category_description
FROM catalog.category c
         JOIN catalog.category_link cl ON c.category_id = cl.linked_category_id
WHERE cl.main_category_id = $1
`

func (q *Queries) FindAllChildrenCategories(ctx context.Context, mainCategoryID pgtype.UUID) ([]CatalogCategory, error) {
	rows, err := q.db.Query(ctx, findAllChildrenCategories, mainCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CatalogCategory{}
	for rows.Next() {
		var i CatalogCategory
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllMainCategories = `-- name: FindAllMainCategories :many
SELECT category_id, category_name, category_description
from catalog.category
WHERE category_id NOT IN (SELECT linked_category_id FROM catalog.category_link)
`

func (q *Queries) FindAllMainCategories(ctx context.Context) ([]CatalogCategory, error) {
	rows, err := q.db.Query(ctx, findAllMainCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CatalogCategory{}
	for rows.Next() {
		var i CatalogCategory
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAllParentCategories = `-- name: FindAllParentCategories :many
SELECT category_id, category_name, category_description
FROM catalog.category c
         JOIN catalog.category_link cl ON c.category_id = cl.main_category_id
WHERE cl.linked_category_id = $1
`

func (q *Queries) FindAllParentCategories(ctx context.Context, linkedCategoryID pgtype.UUID) ([]CatalogCategory, error) {
	rows, err := q.db.Query(ctx, findAllParentCategories, linkedCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CatalogCategory{}
	for rows.Next() {
		var i CatalogCategory
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findCategoryById = `-- name: FindCategoryById :one
SELECT category_id, category_name, category_description
FROM catalog.category
WHERE category_id = $1
LIMIT 1
`

func (q *Queries) FindCategoryById(ctx context.Context, categoryID pgtype.UUID) (CatalogCategory, error) {
	row := q.db.QueryRow(ctx, findCategoryById, categoryID)
	var i CatalogCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription)
	return i, err
}

const findCategoryByIdForUpdate = `-- name: FindCategoryByIdForUpdate :one
SELECT category_id, category_name, category_description
FROM catalog.category
WHERE category_id = $1
LIMIT 1 FOR NO KEY UPDATE
`

func (q *Queries) FindCategoryByIdForUpdate(ctx context.Context, categoryID pgtype.UUID) (CatalogCategory, error) {
	row := q.db.QueryRow(ctx, findCategoryByIdForUpdate, categoryID)
	var i CatalogCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE catalog.category
SET category_name=$2,
    category_description=$3
WHERE category_id = $1
RETURNING category_id, category_name, category_description
`

type UpdateCategoryParams struct {
	CategoryID          pgtype.UUID `json:"category_id"`
	CategoryName        pgtype.Text `json:"category_name"`
	CategoryDescription pgtype.Text `json:"category_description"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (CatalogCategory, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.CategoryID, arg.CategoryName, arg.CategoryDescription)
	var i CatalogCategory
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.CategoryDescription)
	return i, err
}
