package main

type User struct {
	ID       int64  `bun:"id"`
	Name     string `bun:"name"`
	Age      int    `bun:"age"`
	Password string `bun:"password"`
}

func (u User) Model() any {
	return (*User)(nil)
}
