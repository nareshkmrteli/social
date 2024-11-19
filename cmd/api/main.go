package main

import (
	"log"

	db "github.com/nareshkmrteli/social/internal/db"
	environment "github.com/nareshkmrteli/social/internal/env"
	store "github.com/nareshkmrteli/social/internal/store"
)

func main() {
	_cfg := config{
		addr: ":8080",
	}
	_dbConfig := dbConfig{
		addr:         environment.GetString("ADDR", "postgres://admin:adminpassword@localhost/social/sslmode=disabled"),
		maxOpenConns: environment.GetInt("DB_MAX_CONNS", 5),
		maxIdleConns: environment.GetInt("DB_MAX_IDLE_CONNS", 5),
		maxIdleTime:  environment.GetString("DB_MAX_IDLE_TINE", "15min"),
	}
	db_conn, err := db.New(_dbConfig.addr, _dbConfig.maxOpenConns, _dbConfig.maxIdleConns, _dbConfig.maxIdleTime)
	defer db_conn.Close()
	log.Print("closing connection")
	if err != nil {
		log.Fatal(err)
	}
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
