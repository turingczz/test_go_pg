package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

type TGPGDB struct {
	*pg.DB
}

func NewTGPGDB(db *pg.DB) *TGPGDB {
	return &TGPGDB{db}
}

type TGPGDBOptions struct {
	Url string

	DisableBeforeQueryLog bool
	DisableAfterQueryLog  bool
}

func CreateTGPGDB(url string) *TGPGDB {
	return CreateTGPGDBWithOptions(&TGPGDBOptions{Url: url})
}

func CreateTGPGDBWithOptions(dbOpts *TGPGDBOptions) *TGPGDB {
	opts, err := pg.ParseURL(dbOpts.Url)
	if err != nil {
		fmt.Println(DBURLParseError, zap.String("URL", dbOpts.Url), zap.Error(err))
		return nil
	}
	opts.ReadTimeout = DBReadTimeout
	opts.WriteTimeout = DBWriteTimeout
	opts.TLSConfig = nil // disabled for faster local connection (even in production)
	if DBNumConns > 0 {
		opts.PoolSize = DBNumConns
	}

	db := NewTGPGDB(pg.Connect(opts))
	//db.AddQueryHook(dbQueryLoggerHook{
	//	beforeEnabled: !dbOpts.DisableBeforeQueryLog,
	//	afterEnabled:  !dbOpts.DisableAfterQueryLog,
	//})
	return db
}

// Ping simulates a "blank query" behavior similar to lib/pq's
// to check if the db connection is alive.
func (db *TGPGDB) Ping() error {
	_, err := db.ExecOne("SELECT 1")
	return err
}

// Migrate check and migrate to lastest db version.
func (db *TGPGDB) Migrate() error {
	// Make sure to only search specified migrations dir
	cl := migrations.NewCollection()
	cl.DisableSQLAutodiscover(true)
	err := cl.DiscoverSQLMigrations(DBMigrationsDir)
	if err != nil {
		return err
	}

	var oldVersion, newVersion int64
	// Run all migrations in a transaction so we rollback if migrations fail anywhere
	err = db.RunInTransaction(func(tx *pg.Tx) error {
		// Intentionally ignore harmless errors on initializing gopg_migrations
		_, _, err = cl.Run(db, "init")
		//if err != nil && !DBMigrationsAlreadyInit(err) {
		if err != nil{
			return err
		}
		oldVersion, newVersion, err = cl.Run(db, "up")
		return err
	})
	if err != nil {
		return err
	}
	if newVersion == oldVersion {
		fmt.Println("db schema up to date")
	} else {
		fmt.Println("db schema migrated successfully", zap.Int64("from", oldVersion), zap.Int64("to", newVersion))
	}
	return nil
}

// WithContextTimeout executes statements with a default timeout
func WithContextTimeout(ctx context.Context, f func(ctx context.Context)) {
	WithContextTimeoutValue(ctx, DBStmtTimeout, f)
}

// WithContextTimeoutValue executes an inner routine while passing a ctx that supports custom
// timeout and query cancellation to the postgres server.
func WithContextTimeoutValue(ctx context.Context, timeout time.Duration, f func(ctx context.Context)) {
	// check context timeout setting with upper bound read/write limit
	if timeout > DBReadTimeout && timeout > DBWriteTimeout {
		fmt.Println(DBContextTimeoutExceedUpperBound,
			zap.Error(fmt.Errorf("query timeout %s exceed upper bound (%s|%s)", timeout, DBReadTimeout, DBWriteTimeout)))
	}
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	f(newCtx)
	cancel()
}
