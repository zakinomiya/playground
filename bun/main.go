package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func main() {
	engine, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:23306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(engine, mysqldialect.New())

	tables := []Table{User{}}
	if err := createTables(context.Background(), db, tables); err != nil {
		log.Fatalln()
	}
}

type Table interface {
	Model() any
}

func createTables(ctx context.Context, db *bun.DB, tables []Table) error {
	return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		for _, t := range tables {
			_, err := db.NewCreateTable().Model(t.Model()).Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
