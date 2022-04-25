package main

func main() {
	// establish the database conn.
	// dbi := db.NewDBInstance()
	// instantiate and run new app
	// app := application.NewApp(dbi)
	event := InitializeEvent()
	event.RunApp()
	// app.RunApp()
}
