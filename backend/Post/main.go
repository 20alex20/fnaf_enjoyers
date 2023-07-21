package main

import (
	_ "github.com/fnaf-enjoyers/post-service/docs"
	"github.com/fnaf-enjoyers/post-service/pkg/cmd"
	_ "github.com/jackc/pgx"
)

func main() {
	cmd.Run()
}
