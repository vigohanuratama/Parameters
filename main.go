package main

import (
	"Parameters/src/infrastructure"
	"github.com/labstack/echo"
)

var (
	db             *infrastructure.DbHandler
	err            error
	postgresConfig = infrastructure.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "database_manajemen_keuangan",
	}
)

func main() {
	e := echo.New()

	group := e.Group("/api")

	db, err := infrastructure.NewDbHandler(&postgresConfig)

	if err != nil {
		panic(err.Error())
	}

	infrastructure.Init(group, db)

	e.Logger.Fatal(e.Start(":" + "3000"))
}
