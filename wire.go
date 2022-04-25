//go:build wireinject

package main

import (
	"github.com/Siriayanur/GoConcurrency/application"
	"github.com/Siriayanur/GoConcurrency/db"
	"github.com/google/wire"
)

func InitializeEvent() application.IApp {
	wire.Build(application.NewApp, db.NewDBInstance)
	return &application.App{}
}
