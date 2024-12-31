package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/panjiasmoroart/gopher-social/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
