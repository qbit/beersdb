-- name: CreateBrewery :one
INSERT INTO bdb_breweries (
	name, url, location
) VALUES (
	$1, $2, $3
) RETURNING brewery_id, created_at;

-- name: CreateType :one
INSERT INTO bdb_types (
	name
) VALUES (
	$1
) RETURNING type_id, created_at;

-- name: CreateBeer :one
INSERT INTO bdb_beers (
	brewery_id, type_id, name, description, abv, ibu
) VALUES (
	$1, $2, $3, $4, $5, $6
) RETURNING beer_id, created_at;

-- name: CreateUser :one
INSERT INTO bdb_users (
	first_name, last_name, username, email, hash
) VALUES (
	$1, $2, $3, $4, hash($5)
) RETURNING user_id, username, token, token_expires;

-- name: GetUserByToken :one
SELECT * FROM bdb_users
WHERE token = $1 LIMIT 1;
