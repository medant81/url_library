package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/medant81/url_library/internal/config"
	"github.com/medant81/url_library/utils"
	"log/slog"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	//BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	//BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig, log *slog.Logger) (pool *pgxpool.Pool, err error) {

	log.Debug("Start client pgx connection")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	log.Debug("Dsn: ", dsn)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			log.Debug("Error connect: ", err)
			return err
		}
		log.Debug("Connect is ready: ", pool)
		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Error("Error do with tries postgresql: ", err)
	}

	return pool, nil
}
