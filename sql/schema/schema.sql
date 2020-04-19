CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS bdb_beers;
DROP TABLE IF EXISTS bdb_users;
DROP TABLE IF EXISTS bdb_countries;
DROP TABLE IF EXISTS bdb_types;
DROP TABLE IF EXISTS bdb_breweries;

CREATE TABLE bdb_users (
	user_id		BIGSERIAL	NOT NULL PRIMARY KEY UNIQUE,
	created_at	timestamp 	NOT NULL DEFAULT NOW(),
	updated_at	timestamp,
	active		boolean		NOT NULL DEFAULT false,
	first_name	text		NOT NULL,
	last_name	text		NOT NULL,
	username	text 		NOT NULL UNIQUE,
	hash		text 		NOT NULL,
	email		text		NOT NULL,
	token		text		NOT NULL default encode(digest(gen_random_uuid()::text || now(), 'sha256'), 'hex') UNIQUE,
	token_expires	timestamp	NOT NULL DEFAULT NOW() + INTERVAL '30 days'
);

CREATE TABLE bdb_countries (
	id	VARCHAR	NOT NULL PRIMARY KEY UNIQUE,
	value	text	NOT NULL
);

CREATE TABLE bdb_breweries (
	brewery_id	BIGSERIAL	NOT NULL PRIMARY KEY UNIQUE,
	created_at	timestamp 	NOT NULL DEFAULT NOW(),
	name		text 		NOT NULL,
	description	text 		NOT NULL,
	address		text,
	city		text,
	state		text,
	country_id	varchar		NOT NULL REFERENCES bdb_countries (id),
	phone		text,
	url		text
);

CREATE TABLE bdb_types (
	type_id		BIGSERIAL	PRIMARY KEY,
	created_at	timestamp 	NOT NULL DEFAULT NOW(),
	name		text		NOT NULL
);

CREATE TABLE bdb_beers (
	beer_id		BIGSERIAL	NOT NULL PRIMARY KEY UNIQUE,
	brewery_id	BIGSERIAL	NOT NULL REFERENCES bdb_breweries ON DELETE CASCADE,
	type_id		BIGSERIAL	NOT NULL,
	created_at	timestamp 	NOT NULL DEFAULT NOW(),
	updated_at	timestamp,
	name		text		NOT NULL DEFAULT '',
	description	text		NOT NULL DEFAULT '',
	abv		float		NOT NULL,
	ibu		int		NOT NULL
);

CREATE INDEX beer_trgm_idx ON bdb_beers USING gist (description gist_trgm_ops);

CREATE OR REPLACE FUNCTION hash(password text) RETURNS text AS $$
	SELECT crypt(password, gen_salt('bf', 10));
$$ LANGUAGE SQL;
