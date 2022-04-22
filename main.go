package main

import (
	"github.com/Siriayanur/GoConcurrency/application"
	"github.com/Siriayanur/GoConcurrency/db"
)

func main() {
	// establish the database conn.
	dbi := db.NewDBInstance()
	// instantiate and run new app
	app := application.NewApp(dbi)
	app.RunApp()
}
