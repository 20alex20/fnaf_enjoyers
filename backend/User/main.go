package main

import (
	"github.com/fnaf-enjoyers/user/app"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	app.Run()
}
