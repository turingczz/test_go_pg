package main

import (
	"fmt"
	"time"

	pg "github.com/test_go_pg/pg"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting go-pg-migrations...")

	// Bootstrap check pg
	if err := pg.PGDBWrite.Ping(); err != nil {
		fmt.Println(pg.DBWriteConnectionError, zap.Error(err))
		return
	}
	fmt.Println("PostgreSQL is running",
		zap.String("user", pg.PGDBWrite.Options().User),
		zap.String("addr", pg.PGDBWrite.Options().Addr),
		zap.String("db", pg.PGDBWrite.Options().Database))

	// Migrate to latest pg schema
	if err := pg.PGDBWrite.Migrate(); err != nil {
		fmt.Println(pg.DBMigrationError, zap.Error(err))
		return
	}

	time.Sleep(time.Hour)
	fmt.Println("go-pg-migrations is stopping...")
}
