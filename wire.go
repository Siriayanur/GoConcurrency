//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Siriayanur/GoConcurrency/application"
	"github.com/Siriayanur/GoConcurrency/db"
	"github.com/google/wire"
)

func InitializeEvent() *application.App {

	wire.Build(
		wire.NewSet(wire.InterfaceValue(new(db.IDB), db.NewDBInstance()),
			application.NewApp,
		))
	return &application.App{}
}
