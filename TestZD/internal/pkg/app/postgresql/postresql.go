package postgresql

import (
	"TestZD/internal/pkg/utils"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

//type StorageConfig struct {
//username, password, host, port, database string
//	maxAttempts int
//}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context) (pool *pgxpool.Pool, err error) {
	//dsn := fmt.Sprintf("posgresql://%s:%s@%s:%s/%s", sc.username, sc.password, sc.host, sc.port, sc.database)
	err = utils.DoWithTries(func() error {
		ctx, cancle := context.WithTimeout(ctx, 5*time.Second)
		defer cancle()
		//		pool, err = pgxpool.Connect(ctx, "dsn")
		pool, err = pgxpool.New(ctx, "dsn")
		if err != nil {
			return err
		}
		return nil
	}, 5, 5*time.Second)

	return pool, nil
}
