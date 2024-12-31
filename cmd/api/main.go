package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/panjiasmoroart/gopher-social/internal/env"
	"github.com/panjiasmoroart/gopher-social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
