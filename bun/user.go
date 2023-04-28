package main

import (
	"context"
	"math/rand"

	"github.com/google/uuid"
)

type User struct {
	ID       int64  `bun:",pk,autoincrement"`
	Name     string `bun:"name"`
	Age      int    `bun:"age"`
	Password string `bun:"password"`
}

func (u User) GetModel() any {
	return (*User)(nil)
}

func newRandomUser() *User {
	uid := uuid.NewString()
	return &User{
		Name:     uid,
		Age:      rand.Intn(70),
		Password: uid,
	}
}

func (u *User) save(ctx context.Context, db DB) error {
	_, err := db.InsertUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
