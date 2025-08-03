package database

import (
	"context"
	"fmt"

	"github.com/Kenasvarghese/Booking-App/Backend/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbConfig struct {
	config string
}
type db struct {
	pool *pgxpool.Pool
}
type DB interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
}

func LoadDB(cfg *config.Config) DB {
	var newDBConfig dbConfig
	newDBConfig.config = fmt.Sprintf(`host=%s port=%d dbname=%s password=%s user=%s sslmode=%s search_path=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_idle_time=%s`, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBPassword, cfg.DBUser, cfg.SSLMode, cfg.SearchPath, cfg.PoolMaxConns, cfg.PoolMinConns, cfg.PoolMaxConnIdleTime)
	db, err := newDBConfig.NewDBConnection(cfg)
	if err != nil {
		panic(fmt.Errorf("db connection failed: %w",err))
	}
	return db
}
func (dbConfig *dbConfig) NewDBConnection(cfg *config.Config) (DB, error) {
	pgxConfig, err := pgxpool.ParseConfig(dbConfig.config)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &db{pool: pool}, nil

}

func (db *db) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, sql, arguments...)
}

func (db *db) Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, sql, arguments...)
}

func (db *db) QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row {
	return db.pool.QueryRow(ctx, sql, arguments...)
}
