package main

func main() {
	cfg := config{
		addr: ":8080",
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	app.run(mux)

}
