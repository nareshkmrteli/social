package main

import (
	store "./../../internal/store"
)

func main() {
	cfg := config{
		addr: ":8080",
	}
	store := store.Storage()
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	app.run(mux)

}
