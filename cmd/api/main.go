package main

import (
	"log"

	db "github.com/nareshkmrteli/social/internal/db"
	environment "github.com/nareshkmrteli/social/internal/env"
	store "github.com/nareshkmrteli/social/internal/store"
)

func main() {
	_cfg := config{
		addr:    ":8080",
		env:     "dev",
		version: "0.1.0",
	}
	_dbConfig := dbConfig{
		addr:         environment.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
		maxOpenConns: environment.GetInt("DB_MAX_CONNS", 5),
		maxIdleConns: environment.GetInt("DB_MAX_IDLE_CONNS", 5),
		maxIdleTime:  environment.GetString("DB_MAX_IDLE_TIME", "15m"),
	}

	db_conn, err := db.New(_dbConfig.addr, _dbConfig.maxOpenConns, _dbConfig.maxIdleConns, _dbConfig.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db_conn.Close()
	log.Print("closing connection")
	log.Printf("connection open with params %v", _dbConfig)

	store := store.PostgresStorage(db_conn)
	app := &application{
		config: _cfg,
		store:  store,
		db:     _dbConfig,
	}
	mux := app.mount()
	app.run(mux)

}
