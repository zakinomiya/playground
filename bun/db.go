package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Model interface {
	GetModel() any
}

type DB interface {
	CreateTables(ctx context.Context, tables []Model) error
	InsertUser(ctx context.Context, user *User) (*User, error)
}

type BunDB struct {
	*bun.DB
}

func newDB() DB {
	engine, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:23306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	return BunDB{bun.NewDB(engine, mysqldialect.New())}
}

func (bdb BunDB) CreateTables(ctx context.Context, tables []Model) error {
	return bdb.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		for _, t := range tables {
			_, err := bdb.NewCreateTable().Model(t.GetModel()).Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (bdb BunDB) InsertUser(ctx context.Context, user *User) (*User, error) {
	_, err := bdb.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return user, err
	}

	return user, nil
}

func insert[T Model](ctx context.Context, bdb *bun.DB, data T) (T, error) {
	_, err := bdb.NewInsert().Model(data).Exec(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}
