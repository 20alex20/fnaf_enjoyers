package main

import (
	"github.com/fnaf-enjoyers/user-service/app"
	_ "github.com/fnaf-enjoyers/user-service/docs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app.Run()
}
