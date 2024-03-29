-- name: CreateBrewery :one
INSERT INTO bdb_breweries (
	name, url, description, address, city, state, country_id, phone, url
) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING brewery_id, created_at;

-- name: CreateType :one
INSERT INTO bdb_types (
	name,
	description
) VALUES (
	$1, $2
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

-- name: GenerateNewToken :one
UPDATE bdb_users set
	token = DEFAULT,
	token_expires = DEFAULT
WHERE
	token = $1
RETURNING token, token_expires;

-- name: Login :one
SELECT token, token_expires FROM bdb_users
WHERE username = $1 and crypt($2, hash) = hash;

-- name: GetUserByToken :one
SELECT * FROM bdb_users
WHERE token = $1 LIMIT 1;

-- name: GetBeersByBrewery :many
SELECT * FROM bdb_beers
WHERE brewery_id = $1;

-- name: GetRecentBeers :many
SELECT * FROM bdb_beers
WHERE created_at >= $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetAllBeers :many
SELECT * FROM bdb_beers
ORDER BY created_at DESC ;

-- name: SearchBeers :many
SELECT beer_id, brewery_id, name,
	similarity(description, $1) as desc_similarity,
	ts_headline('english', description, q, 'StartSel = <b>, StopSel = </b>') as headline
FROM bdb_beers, to_tsquery($1) q
WHERE similarity(description, $1) > 0.0
	order by similarity DESC;
