// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createBeer = `-- name: CreateBeer :one
INSERT INTO bdb_beers (
	brewery_id, type_id, name, description, abv, ibu
) VALUES (
	$1, $2, $3, $4, $5, $6
) RETURNING beer_id, created_at
`

type CreateBeerParams struct {
	BreweryID   int64   `json:"brewery_id"`
	TypeID      int64   `json:"type_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Abv         float64 `json:"abv"`
	Ibu         int32   `json:"ibu"`
}

type CreateBeerRow struct {
	BeerID    int64     `json:"beer_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateBeer(ctx context.Context, arg CreateBeerParams) (CreateBeerRow, error) {
	row := q.queryRow(ctx, q.createBeerStmt, createBeer,
		arg.BreweryID,
		arg.TypeID,
		arg.Name,
		arg.Description,
		arg.Abv,
		arg.Ibu,
	)
	var i CreateBeerRow
	err := row.Scan(&i.BeerID, &i.CreatedAt)
	return i, err
}

const createBrewery = `-- name: CreateBrewery :one
INSERT INTO bdb_breweries (
	name, url, location
) VALUES (
	$1, $2, $3
) RETURNING brewery_id, created_at
`

type CreateBreweryParams struct {
	Name     string         `json:"name"`
	Url      sql.NullString `json:"url"`
	Location sql.NullString `json:"location"`
}

type CreateBreweryRow struct {
	BreweryID int64     `json:"brewery_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateBrewery(ctx context.Context, arg CreateBreweryParams) (CreateBreweryRow, error) {
	row := q.queryRow(ctx, q.createBreweryStmt, createBrewery, arg.Name, arg.Url, arg.Location)
	var i CreateBreweryRow
	err := row.Scan(&i.BreweryID, &i.CreatedAt)
	return i, err
}

const createType = `-- name: CreateType :one
INSERT INTO bdb_types (
	name
) VALUES (
	$1
) RETURNING type_id, created_at
`

type CreateTypeRow struct {
	TypeID    int64     `json:"type_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateType(ctx context.Context, name string) (CreateTypeRow, error) {
	row := q.queryRow(ctx, q.createTypeStmt, createType, name)
	var i CreateTypeRow
	err := row.Scan(&i.TypeID, &i.CreatedAt)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO bdb_users (
	first_name, last_name, username, email, hash
) VALUES (
	$1, $2, $3, $4, hash($5)
) RETURNING user_id, username, token, token_expires
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserRow struct {
	UserID       int64     `json:"user_id"`
	Username     string    `json:"username"`
	Token        string    `json:"token"`
	TokenExpires time.Time `json:"token_expires"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Token,
		&i.TokenExpires,
	)
	return i, err
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT user_id, created_at, updated_at, first_name, last_name, username, hash, email, token, token_expires FROM bdb_users
WHERE token = $1 LIMIT 1
`

func (q *Queries) GetUserByToken(ctx context.Context, token string) (BdbUser, error) {
	row := q.queryRow(ctx, q.getUserByTokenStmt, getUserByToken, token)
	var i BdbUser
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Hash,
		&i.Email,
		&i.Token,
		&i.TokenExpires,
	)
	return i, err
}
