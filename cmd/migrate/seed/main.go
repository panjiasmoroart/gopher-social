package main

import (
	"log"

	"github.com/panjiasmoroart/gopher-social/internal/db"
	"github.com/panjiasmoroart/gopher-social/internal/env"
	"github.com/panjiasmoroart/gopher-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://new_userpsql:new_userpsql@localhost/go_gopher_social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store, conn)

}
