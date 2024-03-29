// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createBeerStmt, err = db.PrepareContext(ctx, createBeer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBeer: %w", err)
	}
	if q.createBreweryStmt, err = db.PrepareContext(ctx, createBrewery); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBrewery: %w", err)
	}
	if q.createTypeStmt, err = db.PrepareContext(ctx, createType); err != nil {
		return nil, fmt.Errorf("error preparing query CreateType: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.generateNewTokenStmt, err = db.PrepareContext(ctx, generateNewToken); err != nil {
		return nil, fmt.Errorf("error preparing query GenerateNewToken: %w", err)
	}
	if q.getAllBeersStmt, err = db.PrepareContext(ctx, getAllBeers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllBeers: %w", err)
	}
	if q.getBeersByBreweryStmt, err = db.PrepareContext(ctx, getBeersByBrewery); err != nil {
		return nil, fmt.Errorf("error preparing query GetBeersByBrewery: %w", err)
	}
	if q.getRecentBeersStmt, err = db.PrepareContext(ctx, getRecentBeers); err != nil {
		return nil, fmt.Errorf("error preparing query GetRecentBeers: %w", err)
	}
	if q.getUserByTokenStmt, err = db.PrepareContext(ctx, getUserByToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByToken: %w", err)
	}
	if q.loginStmt, err = db.PrepareContext(ctx, login); err != nil {
		return nil, fmt.Errorf("error preparing query Login: %w", err)
	}
	if q.searchBeersStmt, err = db.PrepareContext(ctx, searchBeers); err != nil {
		return nil, fmt.Errorf("error preparing query SearchBeers: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createBeerStmt != nil {
		if cerr := q.createBeerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBeerStmt: %w", cerr)
		}
	}
	if q.createBreweryStmt != nil {
		if cerr := q.createBreweryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBreweryStmt: %w", cerr)
		}
	}
	if q.createTypeStmt != nil {
		if cerr := q.createTypeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTypeStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.generateNewTokenStmt != nil {
		if cerr := q.generateNewTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing generateNewTokenStmt: %w", cerr)
		}
	}
	if q.getAllBeersStmt != nil {
		if cerr := q.getAllBeersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllBeersStmt: %w", cerr)
		}
	}
	if q.getBeersByBreweryStmt != nil {
		if cerr := q.getBeersByBreweryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBeersByBreweryStmt: %w", cerr)
		}
	}
	if q.getRecentBeersStmt != nil {
		if cerr := q.getRecentBeersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRecentBeersStmt: %w", cerr)
		}
	}
	if q.getUserByTokenStmt != nil {
		if cerr := q.getUserByTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByTokenStmt: %w", cerr)
		}
	}
	if q.loginStmt != nil {
		if cerr := q.loginStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing loginStmt: %w", cerr)
		}
	}
	if q.searchBeersStmt != nil {
		if cerr := q.searchBeersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing searchBeersStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	createBeerStmt        *sql.Stmt
	createBreweryStmt     *sql.Stmt
	createTypeStmt        *sql.Stmt
	createUserStmt        *sql.Stmt
	generateNewTokenStmt  *sql.Stmt
	getAllBeersStmt       *sql.Stmt
	getBeersByBreweryStmt *sql.Stmt
	getRecentBeersStmt    *sql.Stmt
	getUserByTokenStmt    *sql.Stmt
	loginStmt             *sql.Stmt
	searchBeersStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		createBeerStmt:        q.createBeerStmt,
		createBreweryStmt:     q.createBreweryStmt,
		createTypeStmt:        q.createTypeStmt,
		createUserStmt:        q.createUserStmt,
		generateNewTokenStmt:  q.generateNewTokenStmt,
		getAllBeersStmt:       q.getAllBeersStmt,
		getBeersByBreweryStmt: q.getBeersByBreweryStmt,
		getRecentBeersStmt:    q.getRecentBeersStmt,
		getUserByTokenStmt:    q.getUserByTokenStmt,
		loginStmt:             q.loginStmt,
		searchBeersStmt:       q.searchBeersStmt,
	}
}
